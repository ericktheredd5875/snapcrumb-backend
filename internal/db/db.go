package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB Instance
var DB *sql.DB

// Initialize DB
func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}

	log.Println("✅ Successfully connected to the database")
}

// InsertURL: Insert a new URL into the database
func InsertURL(originalURL, shortcode string) error {
	query := `
		INSERT INTO urls (original_url, shortcode)
		VALUES ($1, $2)
	`
	_, err := DB.Exec(query, originalURL, shortcode)
	return err
}
