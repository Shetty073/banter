package jwt

import (
	"fmt"
	"banter/utils/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a new JWT token with the provided claims
func GenerateToken(user_id string, expiration int) (string, error) {
	// Ensure the secret key is available
	secretKey := config.Configs.Jwt.Secret
	if secretKey == "" {
		return "", fmt.Errorf("secret key for JWT is missing")
	}

	// Create a new token with the signing method
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Set expiration time for the token
	expiry := time.Hour * time.Duration(expiration)
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(expiry).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
