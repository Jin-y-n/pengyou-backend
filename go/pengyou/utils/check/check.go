package check

import (
	"regexp"
	"unicode"
)

// Regular expression pattern for name validation
// Names may consist of letters (both uppercase and lowercase), digits, underscores, hyphens, and spaces, with a length between 4 and 20 characters.
var nameRegex = regexp.MustCompile(`^[\w- ]{4,20}$`)

// Regular expression pattern for password strength requirements
// Must include at least one letter and one digit, with a length between 8 and 20 characters.
var passwordRegex = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,20}$`)

// Regular expression pattern for phone number validation, applicable only to Chinese mobile phone numbers.
var phoneRegex = regexp.MustCompile(`^1[345789]\\d{9}$`)

// Regular expression pattern for email address validation.
var emailRegex = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)

// Validates whether the provided password meets the requirements.
func checkPassword(password string) bool {
	if password == "" {
		return false
	}

	// Regular expression to check length and allowed characters
	re := regexp.MustCompile(`^[A-Za-z\d]{8,20}$`)
	if !re.MatchString(password) {
		return false
	}

	// Check for at least one letter and one digit
	hasLetter := false
	hasDigit := false
	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		}
		if hasLetter && hasDigit {
			return true
		}
	}
	return false
}

// Validates whether the provided phone number is valid.
func checkPhone(phone string) bool {
	if phone == "" {
		return false
	}
	return phoneRegex.MatchString(phone)
}

// Validates whether the provided email address is valid.
func checkEmail(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// Validates whether the provided name is valid.
func checkName(name string) bool {
	if name == "" {
		return false
	}
	return nameRegex.MatchString(name)
}
