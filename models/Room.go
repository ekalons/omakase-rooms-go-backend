package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Coordinates struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type Room struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Coordinates   Coordinates        `json:"coordinates" bson:"coordinates"`
	Details       string             `json:"details" bson:"details"`
	MichelinStars int                `json:"michelin_stars" bson:"michelin_stars"`
	Name          string             `json:"name" bson:"name"`
	Neighborhood  string             `json:"neighborhood" bson:"neighborhood"`
	Photo         string             `json:"photo" bson:"photo"`
	Price         int                `json:"price" bson:"price"`
	Rating        float64            `json:"rating" bson:"rating"`
	ReviewCount   int                `json:"review_count" bson:"review_count"`
	ServeStyle    string             `json:"serve_style" bson:"serve_style"`
}
