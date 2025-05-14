package main

import (
	"log"
	"os"
	"testing"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

func TestMain(m *testing.M) {
	log.Println("Running Main Test")

	// Load environment variables
	dbURL := utils.RequiredEnv("TEST_DATABASE_URL")

	// Initialize DB
	db.InitDB(dbURL)
	log.Println("DB Initialized and Migrations Run")

	os.Exit(m.Run())
}
