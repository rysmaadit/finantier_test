package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName                  string
	AppPort                  int
	LogLevel                 string
	Environment              string
	PolygonAPIKey            string
	PolygonBaseURL           string
	AuthServiceBaseURL       string
	EncryptionServiceBaseURL string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Println("Failed load .env")
	}

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Println("Failed load .env.testing")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:                  GetString("APP_NAME"),
		AppPort:                  GetInt("APP_PORT"),
		LogLevel:                 GetString("LOG_LEVEL"),
		Environment:              GetString("ENVIRONMENT"),
		PolygonAPIKey:            GetString("POLYGON_API_KEY"),
		PolygonBaseURL:           GetString("POLYGON_BASE_URL"),
		AuthServiceBaseURL:       GetString("AUTH_SERVICE_BASE_URL"),
		EncryptionServiceBaseURL: GetString("ENCRYPTION_SERVICE_BASE_URL"),
	}

	return appConfig
}
