package api

import (
	"encoding/json"
	"net/http"
)

type FieldError struct {
	Error string `json:"error"`
	Field string `json:"field,omitempty"`
}

func WriteJsonError(w http.ResponseWriter, status int, msg string, field string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errResp := FieldError{
		Error: msg,
		Field: field,
	}

	_ = json.NewEncoder(w).Encode(errResp)
}
