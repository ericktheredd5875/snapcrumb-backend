package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetMigrationsPath() (string, error) {

	// path, _ := filepath.Abs(".")
	// log.Printf("Abs Path: %s", path)

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}
	// log.Printf("Working Directory: %s", wd)
	// return "", nil

	if envPath := os.Getenv("MIGRATION_PATH"); envPath != "" {
		// log.Printf("[1]📦 Using migrations from: %s", envPath)
		return "file://" + filepath.ToSlash(envPath), nil
	}

	// wd, err := os.Getwd()
	// if err != nil {
	// 	return "", fmt.Errorf("failed to get working directory: %v", err)
	// }
	// log.Printf("Working Directory: %s", wd)

	rootPath := filepath.Join(wd, "..", "..")
	migrationPath := filepath.Join(rootPath, "db", "migrations")
	migrationPath = filepath.ToSlash(migrationPath)
	// log.Printf("[2]📦 Using migrations from: %s", migrationPath)
	// return fmt.Sprintf("file://%s", migrationPath), nil
	return "file://" + migrationPath, nil
}

func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create postgres driver: %v", err)
	}

	srcUrl, err := GetMigrationsPath()
	if err != nil {
		log.Fatalf("failed to get migrations path: %v", err)
	}
	// log.Printf("📦 Using migrations from: %s", srcUrl)

	m, err := migrate.NewWithDatabaseInstance(
		srcUrl,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	_ = m.Down()
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}
	switch err {
	case nil:
		log.Printf("✅ Migrations applied successfully.")
	case migrate.ErrNoChange:
		log.Printf("🟡 No migrations to apply.")
	default:
		log.Fatalf("❌ Failed to apply migrations: %v", err)
	}
}
