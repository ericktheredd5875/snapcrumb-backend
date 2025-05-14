package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	// err := godotenv.Load("C:/CodeBases/snapcrumb-backend/.env")
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️ .env file not found")
	}

	log.Println("🔑 Environment variables loaded")
}

// ObtainEnv: Get an environment variable with a fallback
func ObtainEnv(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}

// RequiredEnv: Check if required environment variables are set
func RequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	return val
}
