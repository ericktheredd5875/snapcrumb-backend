package utils

import (
	"log"
	"os"
)

// RequiredEnv: Check if required environment variables are set
func RequiredEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	return val
}
