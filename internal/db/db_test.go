package db

import (
	"os"
	"testing"
)

// Connects to test DB
func TestMain(m *testing.M) {

	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		dsn = "postgres://postgres:2b4gp44g6wr607931@localhost:5432/snapcrumb_test?sslmode=disable"
	}
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
