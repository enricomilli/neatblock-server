package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SupabaseClaims struct {
	Sub   string `json:"sub"` // This is the user ID in Supabase
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func ValidJWToken(w http.ResponseWriter, r *http.Request) error {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return errors.New("missing or invalid Authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Get your Supabase JWT secret from environment variable or configuration
	jwtSecret := []byte(os.Getenv("SUPABASE_JWT_SECRET"))

	// Parse and validate the token
	userId, email, err := parseSupabaseToken(token, jwtSecret)
	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	// Add user info to context
	ctx := context.WithValue(r.Context(), "userID", userId)
	ctx = context.WithValue(ctx, "email", email)
	ctx = context.WithValue(ctx, "token", token)

	*r = *r.WithContext(ctx)

	return nil
}

func parseSupabaseToken(tokenString string, jwtSecret []byte) (userId string, email string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &SupabaseClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", "", fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(*SupabaseClaims); ok && token.Valid {
		return claims.Sub, claims.Email, nil
	}

	return "", "", errors.New("invalid token claims")
}
