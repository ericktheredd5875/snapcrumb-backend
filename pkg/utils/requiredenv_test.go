package utils

import (
	"os"
	"testing"
)

func TestRequiredEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	RequiredEnv("TEST_KEY")
}
