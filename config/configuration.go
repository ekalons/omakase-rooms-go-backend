package configuration

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	// MongoDB configuration - support both connection string and individual components
	MongoURL              string // For Railway's MONGO_URL or MONGO_PUBLIC_URL
	MongoDBUsername       string // Fallback to individual components
	MongoDBPassword       string
	MongoDBHost           string
	MongoDBPort           string
	MongoDBCollectionName string
	MongoDBDatabaseName   string
	MongoDBAppName        string
	FrontEndUrl           string
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
		MongoURL:              getMongoURL(),
		MongoDBUsername:       getEnvWithFallback("MONGOUSER", "MONGO_DB_USERNAME"),
		MongoDBPassword:       getEnvWithFallback("MONGOPASSWORD", "MONGO_DB_PASSWORD"),
		MongoDBHost:           getEnvWithFallback("MONGOHOST", "MONGO_DB_HOST"),
		MongoDBPort:           getEnvWithFallback("MONGOPORT", "MONGO_DB_PORT"),
		MongoDBCollectionName: os.Getenv("MONGO_DB_COLLECTION_NAME"),
		MongoDBDatabaseName:   os.Getenv("MONGO_DB_DATABASE_NAME"),
		MongoDBAppName:        os.Getenv("MONGO_DB_APP_NAME"),
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

	// Check MongoDB configuration - either MongoURL or individual components must be set
	if cfg.MongoURL == "" {
		mongoFields := []string{"MongoDBUsername", "MongoDBPassword", "MongoDBHost", "MongoDBPort"}
		for _, fieldName := range mongoFields {
			field, _ := typeOfCfg.FieldByName(fieldName)
			value := v.FieldByName(fieldName)
			if value.Interface() == "" {
				return fmt.Errorf("environment variable for %s is not set (required when MongoURL is not provided)", field.Name)
			}
		}
	}

	// Check other required fields (excluding MongoDB fields which are handled above)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typeOfCfg.Field(i)

		// Skip MongoDB fields as they're handled above
		if fieldType.Name == "MongoURL" || fieldType.Name == "MongoDBUsername" ||
			fieldType.Name == "MongoDBPassword" || fieldType.Name == "MongoDBHost" ||
			fieldType.Name == "MongoDBPort" {
			continue
		}

		if field.Interface() == "" {
			return fmt.Errorf("environment variable for %s is not set", fieldType.Name)
		}
	}

	return nil
}

func getMongoURL() string {
	// Try Railway's connection strings first
	if url := os.Getenv("MONGO_URL"); url != "" {
		return url
	}
	if url := os.Getenv("MONGO_PUBLIC_URL"); url != "" {
		return url
	}
	return ""
}

func getEnvWithFallback(primary, fallback string) string {
	if value := os.Getenv(primary); value != "" {
		return value
	}
	return os.Getenv(fallback)
}
