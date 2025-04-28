package main

import (
	"log"
	"net/http"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/api"
	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
)

func main() {
	// Welcome Message
	http.HandleFunc("/", api.HomeHandler)

	// POST: Shorten URL
	http.HandleFunc("/shorten", api.ShortenURLHandler)

	// GET: Redirect to original URL (shortcode param)
	http.HandleFunc("/{shortcode}", api.RedirectHandler)

	// Initialize DB
	db.InitDB("postgres://postgres:2b4gp44g6wr607931@localhost:5432/snapcrumb?sslmode=disable")

	port := "8080"
	log.Printf("🚀 SnapCrumb server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("🚨 Failed to start server: %v", err)
	}
}
