package middleware

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"os"
	"strings"
)

func ValidateAPIToken(w http.ResponseWriter, r *http.Request) error {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is present and starts with "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return errors.New("missing or invalid Authorization header")
	}

	// Extract the token
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Get the expected token from an environment variable
	expectedToken := os.Getenv("INTERNAL_API_TOKEN")

	// Perform a constant-time comparison of the tokens
	if subtle.ConstantTimeCompare([]byte(token), []byte(expectedToken)) != 1 {
		http.Error(w, "Invalid API token", http.StatusUnauthorized)
		return errors.New("invalid API token")
	}

	// Token is valid
	return nil
}
