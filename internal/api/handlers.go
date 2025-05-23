package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
	"github.com/go-chi/chi/v5"
)

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
	log.Println("📦 SnapCrumb: Received a request to shorten a URL.")

	// Parse the incoming JSON body
	var req utils.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJsonError(w, http.StatusBadRequest, "Invalid request body", "")
		return
	}

	// Validate the input (Make sure the URL is valid)
	if err := utils.ValidateShortenInput(req); err != nil {

		if vErr, ok := err.(utils.ValidationError); ok {
			WriteJsonError(w, http.StatusUnprocessableEntity, vErr.Message, vErr.Field)
		} else {
			WriteJsonError(w, http.StatusBadRequest, err.Error(), "")
		}
		return
	}

	// Generate a unique shortcode (Using Randome String Generator)
	shortcode := utils.GenerateShortCode(6)

	// Store the URL in the database
	if err := db.InsertURL(req.URL, shortcode, req.ExpiresAt); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Failed to store URL in database", "")
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
		WriteJsonError(w, http.StatusBadRequest, "Shortcode missing", "")
		return
	}

	originalURL, expiresAt, err := db.GetOriginalURLByShortcode(shortcode)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Server error", "")
		return
	}

	if !expiresAt.IsZero() && time.Now().After(expiresAt) {
		WriteJsonError(w, http.StatusGone, "URL expired", "")
		return
	}

	log.Printf("Original URL: %s", originalURL)

	if originalURL == "" {
		WriteJsonError(w, http.StatusNotFound, "Shortcode not found", "")
		return
	}

	db.LogVisit(shortcode, r.RemoteAddr, r.UserAgent(), r.Referer())

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	shortcode := chi.URLParam(r, "shortcode")

	count, lastVisit, err := db.GetStats(shortcode)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Server error", "")
		return
	}

	originalURL, expiresAt, err := db.GetOriginalURLByShortcode(shortcode)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, "Server error", "")
		return
	}

	stats := map[string]interface{}{
		"shortcode":     shortcode,
		"visit_count":   count,
		"last_visit_at": lastVisit.Time,
		"expires_at":    expiresAt,
		"original_url":  originalURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
