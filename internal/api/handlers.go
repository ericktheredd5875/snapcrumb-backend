package api

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "👋 Welcome to SnapCrumb! Shorten your links in a snap.")
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📦 SnapCrumb: Received a request to shorten a URL.")
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🔗 SnapCrumb: Redirecting based on shortcode.")
}
