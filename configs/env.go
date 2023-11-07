package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	env := os.Getenv("USE_DOT_ENV")
	// if flag not set load .env file
	if env == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func EnvHost() string {
	loadDotEnv()
	return os.Getenv("HOST")
}

func EnvPort() string {
	loadDotEnv()
	return os.Getenv("PORT")
}

func EnvReviewServicePort() string {
	loadDotEnv()
	return os.Getenv("REVIEW_SERVICE_PORT")
}

func EnvNotificationServicePort() string {
	loadDotEnv()
	return os.Getenv("REVIEW_NOTIFICATION_PORT")
}
