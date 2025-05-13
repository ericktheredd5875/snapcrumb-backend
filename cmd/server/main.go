package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ericktheredd5875/snapcrumb-backend/internal/api"
	"github.com/ericktheredd5875/snapcrumb-backend/internal/db"
	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

func main() {

	// Load environment variables
	port := utils.ObtainEnv("PORT", "8080")
	dbURL := utils.RequiredEnv("DATABASE_URL")

	// Initialize Router
	router := mux.NewRouter()

	// Welcome Message
	router.HandleFunc("/", api.HomeHandler).Methods("GET")

	// POST: Shorten URL
	router.HandleFunc("/shorten", api.ShortenURLHandler).Methods("POST")

	// GET: Redirect to original URL (shortcode param)
	router.HandleFunc("/{shortcode}", api.RedirectHandler).Methods("GET")

	// Initialize DB
	db.InitDB(dbURL)

	log.Printf("🚀 SnapCrumb server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("🚨 Failed to start server: %v", err)
	}
}
