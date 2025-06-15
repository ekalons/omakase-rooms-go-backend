package db

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/ekalons/omakase-rooms-go-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	configuration "github.com/ekalons/omakase-rooms-go-backend/config"
)

var mongoClient *mongo.Client

func Connect() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Temporarily use direct connection in PROD to test basic connectivity
	mongoUri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
		configuration.Cfg.MongoDBUsername,
		configuration.Cfg.MongoDBPassword,
		configuration.Cfg.MongoDBClustername,
		configuration.Cfg.MongoDBAppName,
	)
	fmt.Printf("🔧 %s environment: Testing direct MongoDB connection\n", configuration.Cfg.Environment)

	opts := options.Client().
		ApplyURI(mongoUri).
		SetTLSConfig(&tls.Config{}).
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(5 * time.Minute).
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	mongoClient = client

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return err
	}

	if configuration.Cfg.Environment == "PROD" {
		fmt.Println("✅ Successfully connected to MongoDB directly!")
	}
	return nil
}

func Disconnect() {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	fmt.Println("Disconnecting from MongoDB.")
}

func FetchAllRooms() ([]models.Room, error) {
	var rooms []models.Room

	coll := mongoClient.Database(configuration.Cfg.MongoDBDatabaseName).Collection(configuration.Cfg.MongoDBCollectionName)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var room models.Room
		if err = cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func FetchRoomById(id string) (*models.Room, error) {
	// Convert the string ID to a MongoDB ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	// Create a filter to find the document by its ObjectID
	filter := bson.M{"_id": objectId}

	// Create a variable to hold the result
	var room models.Room

	// Find the document in the collection and decode it into the result variable
	collection := mongoClient.Database(configuration.Cfg.MongoDBDatabaseName).Collection(configuration.Cfg.MongoDBCollectionName)
	err = collection.FindOne(context.TODO(), filter).Decode(&room)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &room, nil
}

func InsertRoom(newRoom models.Room) (*mongo.InsertOneResult, error) {
	if newRoom.ID.IsZero() {
		newRoom.ID = primitive.NewObjectID()
	}

	collection := mongoClient.Database(configuration.Cfg.MongoDBDatabaseName).Collection(configuration.Cfg.MongoDBCollectionName)
	result, err := collection.InsertOne(context.TODO(), newRoom)

	if err != nil {
		return nil, err
	}

	return result, nil
}
