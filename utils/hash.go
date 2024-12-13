package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
	"unicode"
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

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasNumber := false
	hasSymbol := false
	hasSpace := false

	// Iterate through each character in the password
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasNumber = true
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSymbol = true
		}
		if unicode.IsSpace(char) {
			hasSpace = true
		}
	}

	var validationErrors []string

	if hasSpace {
		validationErrors = append(validationErrors, "password must not contain spaces")
	}

	if !hasNumber {
		validationErrors = append(validationErrors, "password must contain at least one number")
	}

	if !hasSymbol {
		validationErrors = append(validationErrors, "password must contain at least one symbol")
	}

	if len(validationErrors) > 0 {
		return errors.New("password validation failed: " + strings.Join(validationErrors, ", "))
	}

	return nil
}

func EncryptAES(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Generate a random IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the data
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptAES(ciphertext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Decode the base64-encoded ciphertext
	data, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// Decrypt the data
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return string(data), nil
}
