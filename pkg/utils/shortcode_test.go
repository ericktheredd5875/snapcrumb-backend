package utils

import (
	"testing"
)

func TestGenerateShortCode(t *testing.T) {

	code := GenerateShortCode(6)
	if len(code) != 6 {
		t.Fatalf("Expected shortcode length to be 6, got %d", len(code))
	}

}
