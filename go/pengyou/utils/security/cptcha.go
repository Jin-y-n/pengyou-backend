package security

import (
	"math/rand"
	"time" // Added to seed the random number generator
)

const numberBytes = "0123456789"
const characterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const defaultLength = 6 // Default captcha length

// SeedRandom seeds the random number generator.
func SeedRandom() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateCaptcha generates a captcha code with the default length.
func GenerateCaptcha() string {
	SeedRandom()

	// Create a byte slice with the default length
	b := make([]byte, defaultLength)
	for i := range b {
		// Randomly select a character from numberBytes
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b)
}

// GenerateCaptchaWithLength generates a captcha code with a specified length.
func GenerateCaptchaWithLength(length int) string {
	SeedRandom()

	// Create a byte slice with the specified length
	b := make([]byte, length)
	for i := range b {
		// Randomly select a character from numberBytes
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}

	return string(b)
}

// GenerateCaptchaWithCharacter generates a captcha code with a specified length and characters.
func GenerateCaptchaWithCharacter(length int) string {
	SeedRandom()

	// Create a byte slice with the specified length
	b := make([]byte, length)
	for i := range b {
		// Randomly select a character from characterBytes
		b[i] = characterBytes[rand.Intn(len(characterBytes))]
	}

	return string(b)
}
