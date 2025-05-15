package db

import (
	"database/sql"
	"log"
	"time"

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

	RunMigrations(DB)
}

// InsertURL: Insert a new URL into the database
//   - Todo: If Insert fails with a unique_violation error,
//     generate a new shortcode and retry.
func InsertURL(originalURL string, shortcode string, expiresAt *time.Time) error {
	query := `
		INSERT INTO urls (original_url, shortcode, expires_at)
		VALUES ($1, $2, $3)
	`
	_, err := DB.Exec(query, originalURL, shortcode, expiresAt)
	return err
}

func GetOriginalURLByShortcode(shortcode string) (string, time.Time, error) {
	var originalURL string
	var expiresAt sql.NullTime
	query := `
		SELECT original_url, expires_at 
			FROM urls WHERE shortcode = $1;
	`

	err := DB.QueryRow(query, shortcode).Scan(&originalURL, &expiresAt)
	if err == sql.ErrNoRows {
		return "", time.Time{}, nil
	}
	if err != nil {
		return "", time.Time{}, err
	}

	return originalURL, expiresAt.Time, nil
}
