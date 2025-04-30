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
//   - Todo: If Insert fails with a unique_violation error,
//     generate a new shortcode and retry.
func InsertURL(originalURL, shortcode string) error {
	query := `
		INSERT INTO urls (original_url, shortcode)
		VALUES ($1, $2)
	`
	_, err := DB.Exec(query, originalURL, shortcode)
	return err
}

func GetOriginalURLByShortcode(shortcode string) (string, error) {
	var originalURL string
	query := "SELECT original_url FROM urls WHERE shortcode = $1;"
	err := DB.QueryRow(query, shortcode).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
