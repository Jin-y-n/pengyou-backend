package security

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// MD5Encrypt returns the hexadecimal representation of the MD5 hash of the given string.
func MD5Encrypt(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA1Encrypt returns the hexadecimal representation of the SHA-1 hash of the given string.
func SHA1Encrypt(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA256Encrypt returns the hexadecimal representation of the SHA-256 hash of the given string.
func SHA256Encrypt(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// SHA512Encrypt returns the hexadecimal representation of the SHA-512 hash of the given string.
func SHA512Encrypt(input string) string {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// BCryptEncrypt returns the bcrypt hash of the given string or an error if hashing fails.
func BCryptEncrypt(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
