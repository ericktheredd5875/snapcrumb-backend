package main

import (
	"os"
	"testing"

	"github.com/ericktheredd5875/snapcrumb-backend/dbtest"
)

func TestMain(m *testing.M) {
	os.Exit(dbtest.Setup(m))
}
