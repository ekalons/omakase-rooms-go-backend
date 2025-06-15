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
	MongoDBAppName        string
	FrontEndUrl           string
	QuotaGuardStaticURL   string
	ClientSecret          string
	JWTSecret             string
	JWTClaimsSubKey       string
	Environment           string
}

var Cfg Config

func Load() {
	fmt.Println("Environment: ", os.Getenv("ENVIRONMENT"))

	if os.Getenv("ENVIRONMENT") != "PROD" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	Cfg = Config{
		MongoDBUsername:       os.Getenv("MONGO_DB_USERNAME"),
		MongoDBPassword:       os.Getenv("MONGO_DB_PASSWORD"),
		MongoDBClustername:    os.Getenv("MONGO_DB_CLUSTER_NAME"),
		MongoDBCollectionName: os.Getenv("MONGO_DB_COLLECTION_NAME"),
		MongoDBDatabaseName:   os.Getenv("MONGO_DB_DATABASE_NAME"),
		MongoDBAppName:        os.Getenv("MONGO_DB_APP_NAME"),
		FrontEndUrl:           os.Getenv("FRONTEND_URL"),
		QuotaGuardStaticURL:   os.Getenv("QUOTA_GUARD_STATIC_URL"),
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
