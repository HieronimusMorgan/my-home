package utils

import (
	"Master_Data/module/domain"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint, clientID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"client_id": clientID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.UserID,
		"client_id": user.ClientID,
		"uuid_key":  user.UUIDKey,
		"role_id":   user.RoleID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken() (string, error) {
	// Create a random 32-byte token
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Encode the token to a base64 string
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func GetClientIDFromToken(tokenString string) (interface{}, error) {
	// Validate and parse the JWT
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Get the user_id from claims
	clientID, exists := claims["client_id"]
	if !exists {
		return nil, errors.New("client_id not found in token claims")
	}

	return clientID, nil
}

func GetUserIDFromToken(tokenString string) (interface{}, error) {
	// Validate and parse the JWT
	token, err := ValidateJWT(tokenString)
	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	// Get the user_id from claims
	clientID, exists := claims["user_id"]
	if !exists {
		return nil, errors.New("user_id not found in token claims")
	}

	return clientID, nil
}

func Float64ToUint(value interface{}) (uint, error) {
	floatVal, ok := value.(float64)
	if !ok {
		return 0, errors.New("value is not a float64")
	}

	if floatVal < 0 {
		return 0, errors.New("cannot convert negative float64 to uint")
	}

	return uint(floatVal), nil
}
