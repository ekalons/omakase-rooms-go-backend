package configuration

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBUsername       string
	MongoDBPassword       string
	MongoDBClustername    string
	MongoDBCollectionName string
	MongoDBDatabaseName   string
	FrontEndUrl           string
}

var Cfg Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg = Config{
		MongoDBUsername:       os.Getenv("MONGO_DB_USERNAME"),
		MongoDBPassword:       os.Getenv("MONGO_DB_PASSWORD"),
		MongoDBClustername:    os.Getenv("MONGO_DB_CLUSTER_NAME"),
		MongoDBCollectionName: os.Getenv("MONGO_DB_COLLECTION_NAME"),
		MongoDBDatabaseName:   os.Getenv("MONGO_DB_DATABASE_NAME"),
		FrontEndUrl:           os.Getenv("FRONTEND_URL"),
	}
	if err := validateEnvVars(Cfg); err != nil {
		log.Fatal(err)
	}

}

func validateEnvVars(cfg Config) error {
	v := reflect.ValueOf(cfg)
	typeOfCfg := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Interface() == "" {
			return fmt.Errorf("environment variable for %s is not set", typeOfCfg.Field(i).Name)
		}
	}

	return nil
}
