package check

import (
	"strings"
	"unicode"
)

// isNilOrEmpty 检查字符串指针是否为 nil 或者字符串内容是否为空
func IsNilOrEmpty(s *string) bool {
	if s == nil || *s == "" {
		return true
	}
	return false
}

// isBlank 检查字符串是否只包含空白字符
func IsBlank(s *string) bool {
	return strings.TrimSpace(*s) == ""
}

// isNilOrBlank 检查字符串指针是否为 nil 或者字符串内容是否只包含空白字符
func IsNilOrBlank(s *string) (bool, error) {
	if s == nil {
		return true, nil
	}
	if strings.TrimSpace(*s) == "" {
		return true, nil
	}
	return false, nil
}

// IsNilOrEmpty checks if a string pointer is nil or the string is empty.
func UIsNilOrEmpty(s *string) bool {
	if s == nil || *s == "" {
		return true
	}
	return false
}

// IsBlank checks if a string contains only whitespace characters.
// It supports all Unicode whitespace characters.
func UIsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsNilOrBlank checks if a string pointer is nil or the string contains only whitespace.
func UIsNilOrBlank(s *string) bool {
	if s == nil {
		return true
	}
	return IsBlank(s)
}
