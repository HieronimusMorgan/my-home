package utils

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateClientID() string {
	randomBytes := make([]byte, 16) // Generate 16 random bytes
	rand.Read(randomBytes)
	randomPart := hex.EncodeToString(randomBytes) // Convert random bytes to hex
	return randomPart
}
