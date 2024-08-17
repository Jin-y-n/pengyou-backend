package check

import (
	"strings"
	"unicode"
)

// IsNilOrEmpty checks if the string pointer is nil or the string content is empty.
func IsNilOrEmpty(s *string) bool {
	if s == nil || *s == "" {
		return true
	}
	return false
}

// IsBlank checks if the string contains only whitespace characters.
func IsBlank(s *string) bool {
	return strings.TrimSpace(*s) == ""
}

// IsNilOrBlank checks if the string pointer is nil or the string contains only whitespace.
func IsNilOrBlank(s *string) bool {
	if s == nil {
		return true
	}
	return IsBlank(s)
}

// UIsBlank checks if a string contains only whitespace characters.
// It supports all Unicode whitespace characters.
func UIsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// UIsNilOrBlank checks if a string pointer is nil or the string contains only whitespace.
func UIsNilOrBlank(s *string) bool {
	if s == nil {
		return true
	}
	return UIsBlank(*s)
}
