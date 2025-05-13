package dbtest

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

var TestDB *sql.DB

func Setup(m *testing.M) int {

	var err error

	// Load environment variables
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := utils.RequiredEnv("DATABASE_URL")

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
