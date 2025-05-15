package db

import (
	"os"
	"testing"
	"time"

	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

// Connects to test DB
func TestMain(m *testing.M) {

	// Load environment variables
	dsn := utils.RequiredEnv("TEST_DATABASE_URL")
	InitDB(dsn)

	code := m.Run()
	os.Exit(code)
}

func TestInsertAndGetUser(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"
	expiresAt := time.Now().Add(time.Hour * 24)

	// Clean URL out for test
	DB.Exec("DELETE FROM urls WHERE shortcode = $1", shortcode)

	// Insert URL
	err := InsertURL(originalURL, shortcode, &expiresAt)
	if err != nil {
		t.Fatalf("Failed to insert URL: %v", err)
	}

	// Get URL
	result, expiresAt, err := GetOriginalURLByShortcode(shortcode)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}

	if result != originalURL {
		t.Fatalf("Expected %s, got %s", originalURL, result)
	}

	if expiresAt.IsZero() {
		t.Fatalf("Expected expiresAt to be set, got %v", expiresAt)
	}
}

func TestInsertUrl(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"
	expiresAt := time.Now().Add(time.Hour * 24)

	// Clean URL out for test
	DB.Exec("DELETE FROM urls WHERE shortcode = $1", shortcode)

	err := InsertURL(originalURL, shortcode, &expiresAt)
	if err != nil {
		t.Fatalf("Failed to insert URL: %v", err)
	}
}

func TestGetUrl(t *testing.T) {
	originalURL := "https://www.google.com"
	shortcode := "test123456"

	// Get URL
	result, expiresAt, err := GetOriginalURLByShortcode(shortcode)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}

	if result != originalURL {
		t.Fatalf("Expected %s, got %s", originalURL, result)
	}

	if expiresAt.IsZero() {
		t.Fatalf("Expected expiresAt to be set, got %v", expiresAt)
	}

}
