package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

func EnvUserServicePort() string {
	loadDotEnv()
	return os.Getenv("USER_SERVICE_PORT")
}

func EnvOrderServicePort() string {
	loadDotEnv()
	return os.Getenv("ORDER_SERVICE_PORT")
}

func EnvNotificationServicePort() string {
	loadDotEnv()
	return os.Getenv("NOTIFICATION_SERVICE_PORT")
}

func EnvMenuServicePort() string {
	loadDotEnv()
	return os.Getenv("MENU_SERVICE_PORT")
}

func EnvMongoURI() string {
	loadDotEnv()
	return os.Getenv("MONGO_URI")
}

func EnvMongoDBName() string {
	loadDotEnv()
	return os.Getenv("MONGO_DATABASE_NAME")
}

func EnvServicePort() string {
	loadDotEnv()
	return os.Getenv("PORT")
}

func EnvFrontendPort() string {
	loadDotEnv()
	return os.Getenv("FRONTEND_PORT")
}

func EnvAccessTokenPK() string {
	loadDotEnv()
	return os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
}
func EnvAccessTokenPublicK() string {
	loadDotEnv()
	return os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
}
func EnvAccessTokenMaxAge() int {
	loadDotEnv()
	maxAgeStr := os.Getenv("ACCESS_TOKEN_MAXAGE")
	maxAge, err := strconv.Atoi(maxAgeStr)
	if err != nil {
		return 0 // Default value
	}
	return maxAge
}
func EnvAccessTokenExpire() time.Duration {
	loadDotEnv()
	expireStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	if expireStr == "" {
		return 0 // Return zero duration if the environment variable is not set
	}

	duration, err := time.ParseDuration(expireStr)
	if err != nil {
		fmt.Println("Error parsing duration")
		return 0 // Return an error if parsing fails
	}

	return duration
}

func EnvReviewServiceHost() string {
	loadDotEnv()
	return os.Getenv("REVIEW_SERVICE_HOST")
}

func EnvUserServiceHost() string {
	loadDotEnv()
	return os.Getenv("USER_SERVICE_HOST")
}

func EnvNotificationServiceHost() string {
	loadDotEnv()
	return os.Getenv("NOTIFICATION_SERVICE_HOST")
}

func EnvOrderServiceHost() string {
	loadDotEnv()
	return os.Getenv("ORDER_SERVICE_HOST")
}

func EnvMenuServiceHost() string {
	loadDotEnv()
	return os.Getenv("MENU_SERVICE_HOST")
}
