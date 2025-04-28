package main

import (
	"log"
	"net/http"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/api"
)

func main() {
	// Welcome Message
	http.HandleFunc("/", api.HomeHandler)

	// POST: Shorten URL
	http.HandleFunc("/shorten", api.ShortenURLHandler)

	// GET: Redirect to original URL (shortcode param)
	http.HandleFunc("/{shortcode}", api.RedirectHandler)

	port := "8080"
	log.Printf("🚀 SnapCrumb server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("🚨 Failed to start server: %v", err)
	}
}
