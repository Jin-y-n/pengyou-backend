package security

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// MD5Encrypt 函数接收一个字符串并返回其MD5加密后的十六进制表示
func MD5Encrypt(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// SHA1Encrypt 函数接收一个字符串并返回其SHA-1加密后的十六进制表示
func SHA1Encrypt(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// SHA256Encrypt 函数接收一个字符串并返回其SHA-256加密后的十六进制表示
func SHA256Encrypt(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// SHA512Encrypt 函数接收一个字符串并返回其SHA-512加密后的十六进制表示
func SHA512Encrypt(input string) string {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// BCryptEncrypt 函数接收一个字符串并返回其BCrypt加密后的结果
func BCryptEncrypt(input string) (string, error) {
	bytes := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
