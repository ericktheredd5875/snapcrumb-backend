package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "👋 Welcome to SnapCrumb! Shorten your links in a snap.")
}

func main() {
	http.HandleFunc("/", HomeHandler)

	port := "8080"
	log.Printf("🚀 SnapCrumb server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("🚨 Failed to start server: %v", err)
	}
}
