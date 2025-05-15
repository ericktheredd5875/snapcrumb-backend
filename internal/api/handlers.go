package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
	"github.com/go-chi/chi/v5"
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
	shortcode := utils.GenerateShortCode(6)

	// Store the URL in the database
	err = db.InsertURL(req.URL, shortcode)
	if err != nil {
		http.Error(w, "Failed to store URL in database", http.StatusInternalServerError)
		return
	}

	// Create the shortened URL
	url := utils.ObtainEnv("URL", "localhost")
	port := utils.ObtainEnv("PORT", "8080")
	shortened := url + ":" + port + "/" + shortcode

	// Return the shortened URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortenResponse{
		Shortcode:    shortcode,
		ShortenedURL: shortened,
	})
}

// RedirectHandler: Redirect to original URL GET /{shortcode}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	shortcode := chi.URLParam(r, "shortcode")
	log.Printf("Shortcode: %s", shortcode)

	if shortcode == "" {
		http.Error(w, "Shortcode missing", http.StatusBadRequest)
		return
	}

	originalURL, err := db.GetOriginalURLByShortcode(shortcode)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Original URL: %s", originalURL)

	if originalURL == "" {
		http.Error(w, "Shortcode not found", http.StatusNotFound)
		return
	}

	db.LogVisit(db.DB, shortcode, r.RemoteAddr, r.UserAgent(), r.Referer())

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	shortcode := chi.URLParam(r, "shortcode")

	count, lastVisit, err := db.GetStats(db.DB, shortcode)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	stats := map[string]interface{}{
		"shortcode":     shortcode,
		"visit_count":   count,
		"last_visit_at": lastVisit.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
