package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
)

// Request Body struct
type shortenRequest struct {
	URL string `json:"url"`
}

// Response Body struct
type shortenResponse struct {
	Shortcode    string `json:"shortcode"`
	ShortenedURL string `json:"shortened_url"`
}

// HomeHandler: Welcome Message GET /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "👋 Welcome to SnapCrumb! Shorten your links in a snap.")
}

// ShortenURLHandler: Shorten a URL POST /shorten
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📦 SnapCrumb: Received a request to shorten a URL.")

	// Parse the incoming JSON body
	var req shortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input (Make sure the URL is valid)
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	// Generate a unique shortcode (Using Randome String Generator)
	shortcode := generateShortCode(6)

	// Store the URL in the database
	err = db.InsertURL(req.URL, shortcode)
	if err != nil {
		http.Error(w, "Failed to store URL in database", http.StatusInternalServerError)
		return
	}

	// Create the shortened URL
	domain := "http://localhost:8080"
	shortened := domain + "/" + shortcode

	// Return the shortened URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortenResponse{
		Shortcode:    shortcode,
		ShortenedURL: shortened,
	})
}

// RedirectHandler: Redirect to original URL GET /{shortcode}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🔗 SnapCrumb: Redirecting based on shortcode.")
}

// generateShortCode: Generate a random shortcode
func generateShortCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
