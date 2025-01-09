package apiutil

import (
	"net/url"
	"strconv"
	"strings"
)

// A general check to see if the string is an URL
func IsValidURL(str string) bool {
	// Trim spaces
	str = strings.TrimSpace(str)

	// Check length constraints
	if len(str) < 10 || len(str) > 2048 { // Most browsers support up to 2048 characters
		return false
	}

	// Parse the URL
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	// Check scheme
	if u.Scheme == "" || !isValidScheme(u.Scheme) {
		return false
	}

	// Check host
	host := u.Hostname()
	if host == "" || !isValidHost(host) {
		return false
	}

	// Check port if specified
	if port := u.Port(); port != "" && !isValidPort(port) {
		return false
	}

	// Check path
	if !isValidPath(u.Path) {
		return false
	}

	return true
}

func isValidScheme(scheme string) bool {
	scheme = strings.ToLower(scheme)
	return scheme == "http" || scheme == "https"
}

func isValidHost(host string) bool {
	// Check for common invalid patterns
	if strings.Contains(host, "..") ||
		strings.Contains(host, "//") ||
		strings.HasPrefix(host, ".") ||
		strings.HasSuffix(host, ".") {
		return false
	}

	// Check host length
	if len(host) > 253 { // Max length per DNS specs
		return false
	}

	// Check each label
	labels := strings.Split(host, ".")
	if len(labels) < 2 { // Must have at least two labels (e.g., example.com)
		return false
	}

	for _, label := range labels {
		if !isValidHostLabel(label) {
			return false
		}
	}

	return true
}

func isValidHostLabel(label string) bool {
	if len(label) == 0 || len(label) > 63 { // Max label length per DNS specs
		return false
	}

	// Must start and end with alphanumeric
	if !isAlphanumeric(rune(label[0])) || !isAlphanumeric(rune(label[len(label)-1])) {
		return false
	}

	// Check each character
	for _, ch := range label {
		if !isAlphanumeric(ch) && ch != '-' {
			return false
		}
	}

	return true
}

func isValidPort(port string) bool {
	// Convert port to integer
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	// Check port range (0-65535)
	return portNum >= 0 && portNum <= 65535
}

func isValidPath(path string) bool {
	// Check path length
	if len(path) > 2048 {
		return false
	}

	// Check for suspicious patterns
	suspiciousPatterns := []string{
		"../", "/..",
		"//",
		"<", ">",
		"'", "\"",
		";",
		"%00", // null byte
	}

	for _, pattern := range suspiciousPatterns {
		if strings.Contains(path, pattern) {
			return false
		}
	}

	return true
}

func isAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
