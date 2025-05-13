package dbtest

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// TestDB is the shared test database connection
var testDB *sql.DB

func SetupDB(m *testing.M) int {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		dsn = "postgres://postgres:2b4gp44g6wr607931@localhost:5432/snapcrumb_test?sslmode=disable"
	}

	var err error
	testDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to test db: %v", err)
	}
	log.Println("Connected to test db")

	if err := testDB.Ping(); err != nil {
		log.Fatalf("Failed to ping test db: %v", err)
	}
	log.Println("Pinged test db")

	if err := resetSchema(); err != nil {
		log.Fatalf("Failed to reset schema: %v", err)
	}
	log.Println("Reset schema")

	return m.Run()
}

// Reset the schema
func resetSchema() error {
	schema := `
	DROP TABLE IF EXISTS urls;
	CREATE TABLE urls (
		id SERIAL PRIMARY KEY,
		original_url VARCHAR(255) NOT NULL,
		shortcode VARCHAR(255) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := testDB.Exec(schema)
	return err
}

func CleanDB(t *testing.T) {
	_, err := testDB.Exec("TRUNCATE TABLE urls RESTART IDENTITY CASCADE;")
	if err != nil {
		t.Fatalf("Failed to truncate test db: %v", err)
	}
}
