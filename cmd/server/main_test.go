package main

import (
	"log"
	"os"
	"testing"

	"github.com/ericktheredd5875/snapcrumb-backend/dbtest"
)

func TestMain(m *testing.M) {
	log.Println("Running Main Test")
	os.Exit(dbtest.Setup(m))
}
