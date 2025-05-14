package main

import (
	"log"
	"net/http"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/api"
	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func main() {

	// Load environment variables
	port := utils.ObtainEnv("PORT", "8080")
	dbURL := utils.RequiredEnv("DATABASE_URL")

	// Initialize Router
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Matched URL:", r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Welcome Message
	r.Get("/", api.HomeHandler)

	// POST: Shorten URL
	r.Post("/shorten", api.ShortenURLHandler)

	// GET: Redirect to original URL (shortcode param)
	r.Get("/{shortcode}", api.RedirectHandler)

	// Initialize DB
	db.InitDB(dbURL)

	log.Printf("🚀 SnapCrumb server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("🚨 Failed to start server: %v", err)
	}
}
