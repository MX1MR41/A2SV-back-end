package Infrastucture

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// Global variable with which tokens will be hashed with and signed on
var jwtSecret = []byte("shhhh... it's a secret")

// GenerateToken creates a new token with claims including username and role
func GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
	})

	// Hash encoded token with the jwtSecret and append the result
	// to the token as a signature
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the token and returns the claims
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	return token, err
}
