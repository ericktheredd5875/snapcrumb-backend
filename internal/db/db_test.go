package db

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

// Connects to test DB
func TestMain(m *testing.M) {

	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("ℹ️ .env file not found")
	}

	dsn := utils.RequiredEnv("DATABASE_URL")
	InitDB(dsn)

	code := m.Run()
	os.Exit(code)
}

func TestInsertAndGetUser(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"

	// Clean URL out for test
	DB.Exec("DELETE FROM urls WHERE shortcode = $1", shortcode)

	// Insert URL
	err := InsertURL(originalURL, shortcode)
	if err != nil {
		t.Fatalf("Failed to insert URL: %v", err)
	}

	// Get URL
	result, err := GetOriginalURLByShortcode(shortcode)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}

	if result != originalURL {
		t.Fatalf("Expected %s, got %s", originalURL, result)
	}
}

func TestInsertUrl(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"

	// Clean URL out for test
	DB.Exec("DELETE FROM urls WHERE shortcode = $1", shortcode)

	err := InsertURL(originalURL, shortcode)
	if err != nil {
		t.Fatalf("Failed to insert URL: %v", err)
	}
}

func TestGetUrl(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"

	// Get URL
	result, err := GetOriginalURLByShortcode(shortcode)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}

	if result != originalURL {
		t.Fatalf("Expected %s, got %s", originalURL, result)
	}

}
