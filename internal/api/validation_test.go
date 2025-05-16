package api

import (
	"testing"
	"time"

	"github.com/ericktheredd5875/snapcrumb-backend/pkg/utils"
)

func TestValidateShortenInput(t *testing.T) {

	now := time.Now()
	past := now.Add(-24 * time.Hour)
	future := now.Add(24 * time.Hour)

	tests := []struct {
		name       string
		input      utils.ShortenRequest
		wantErr    bool
		wantField  string
		wantErrMsg string
	}{
		{
			name:       "missing URL",
			input:      utils.ShortenRequest{},
			wantErr:    true,
			wantField:  "url",
			wantErrMsg: "url is required",
		},
		{
			name:       "invalid scheme",
			input:      utils.ShortenRequest{URL: "ftp://example.com"},
			wantErr:    true,
			wantField:  "url",
			wantErrMsg: "URL must start with http:// or https://",
		},
		{
			name:       "too long",
			input:      utils.ShortenRequest{URL: "https://" + string(make([]byte, 2001))},
			wantErr:    true,
			wantField:  "url",
			wantErrMsg: "URL is too long",
		},
		{
			name:       "expired",
			input:      utils.ShortenRequest{URL: "https://example.com", ExpiresAt: &past},
			wantErr:    true,
			wantField:  "expires_at",
			wantErrMsg: "expires_at must be in the future",
		},
		{
			name:    "valid request",
			input:   utils.ShortenRequest{URL: "https://example.com", ExpiresAt: &future},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.ValidateShortenInput(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}

				vErr, ok := err.(utils.ValidationError)
				if !ok {
					t.Fatalf("expected ValidationError, got %T", err)
				}
				if vErr.Field != tt.wantField {
					t.Errorf("expected field %s, got %s", tt.wantField, vErr.Field)
				}
				if vErr.Message != tt.wantErrMsg {
					t.Errorf("expected message %s, got %s", tt.wantErrMsg, vErr.Message)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			}
		})
	}
}
