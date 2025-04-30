package shortcode

import (
	"math/rand"
	"strings"
	"time"
)

// generateShortCode: Generate a random shortcode
func GenerateShortCode(length int) string {
	// rand.Seed(time.Now().UnixNano())
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rando.Intn(len(chars))])
	}

	return b.String()
}
