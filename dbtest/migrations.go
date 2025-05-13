package dbtest

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var TestDB *sql.DB

func Setup(m *testing.M) int {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:2b4gp44g6wr607931@localhost:5432/snapcrumb_test?sslmode=disable"
	}

	var err error
	TestDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to test db: %v", err)
	}

	if err := runMigrations(TestDB); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	return m.Run()
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %v", err)
	}

	_ = m.Down()
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %v", err)
	}

	return nil
}
