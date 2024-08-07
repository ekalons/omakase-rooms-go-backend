package configuration

import (
	"fmt"
	"log"
	"os"
	"reflect"
)

type Config struct {
	MongoDBUsername       string
	MongoDBPassword       string
	MongoDBClustername    string
	MongoDBCollectionName string
	MongoDBDatabaseName   string
	FrontEndUrl           string
	ClientSecret          string
	JWTSecret             string
	JWTClaimsSubKey       string
	Environment           string
}

var Cfg Config

func Load() {

	Cfg = Config{
		MongoDBUsername:       os.Getenv("MONGO_DB_USERNAME"),
		MongoDBPassword:       os.Getenv("MONGO_DB_PASSWORD"),
		MongoDBClustername:    os.Getenv("MONGO_DB_CLUSTER_NAME"),
		MongoDBCollectionName: os.Getenv("MONGO_DB_COLLECTION_NAME"),
		MongoDBDatabaseName:   os.Getenv("MONGO_DB_DATABASE_NAME"),
		FrontEndUrl:           os.Getenv("FRONTEND_URL"),
		ClientSecret:          os.Getenv("CLIENT_SECRET"),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		JWTClaimsSubKey:       os.Getenv("JWT_CLAIMS_SUB_KEY"),
		Environment:           os.Getenv("ENVIRONMENT"),
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
