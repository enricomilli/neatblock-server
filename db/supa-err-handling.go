package db

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleSupabaseError(err error) (int, string) {
	// Convert error to string for pattern matching
	errStr := err.Error()
	fmt.Println("error msg:", errStr)

	switch {
	case strings.Contains(errStr, "duplicate key"):
		return http.StatusConflict, "That Pool URL already exists"

	case strings.Contains(errStr, "foreign key violation"):
		return http.StatusBadRequest, "Invalid user ID reference"

	case strings.Contains(errStr, "invalid input syntax"):
		return http.StatusBadRequest, "Invalid data format"

	case strings.Contains(errStr, "value too long"):
		return http.StatusBadRequest, "Pool URL exceeds maximum length"

	case strings.Contains(errStr, "not-authenticated"):
		return http.StatusUnauthorized, "Authentication required"

	case strings.Contains(errStr, "permission denied"):
		return http.StatusForbidden, "You don't have permission to perform this action"

	case strings.Contains(errStr, "violates row-level security"):
		return http.StatusForbidden, "You don't have permission to perform this action"

	case strings.Contains(errStr, "timeout"):
		return http.StatusGatewayTimeout, "Database operation timed out"

	case strings.Contains(errStr, "connection refused"):
		return http.StatusServiceUnavailable, "Database connection failed"
	}

	// Default error
	return http.StatusInternalServerError, "An unexpected database error occurred"
}
