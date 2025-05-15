package utils

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

type ShortenRequest struct {
	URL       string     `json:"url"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

var maxURLLength = 2000
var disallowedURLs = []string{
	"localhost",
	"127.0.0.1",
	"file://",
	"data://",
	"javascript:",
	"vbscript:",
	"about:",
	"internal.company",
	"malware.com",
	"phishing.com",
	"example.com",
}

func ValidateShortenInput(req ShortenRequest) error {

	if err := ValidateURL(req.URL); err != nil {
		return err
	}

	if req.ExpiresAt != nil && time.Now().After(*req.ExpiresAt) {
		return errors.New("expires_at cannot be in the past")
	}

	return nil
}

func ValidateURL(srcUrl string) error {

	// Check if URL is empty
	if srcUrl == "" {
		return errors.New("URL is required")
	}

	// Check if URL starts with http:// or https://
	if !strings.HasPrefix(srcUrl, "http://") && !strings.HasPrefix(srcUrl, "https://") {
		return errors.New("URL must start with http:// or https://")
	}

	//  Maximum URL length check
	if len(srcUrl) > maxURLLength {
		return errors.New("URL is too long")
	}

	log.Printf("Length: %d", len(srcUrl))

	// Disallowed URLs check
	for _, disallowed := range disallowedURLs {
		if strings.Contains(srcUrl, disallowed) {
			return fmt.Errorf("url domain '%s' is not allowed", disallowed)
		}
	}

	// Validate URL format
	_, err := url.Parse(srcUrl)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	return nil
}
