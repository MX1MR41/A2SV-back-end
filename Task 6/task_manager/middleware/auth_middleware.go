package middleware

import (
	"fmt"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Global variable with which tokens will be hashed with and signed on
var jwtSecret = []byte("shhhh... it's a secret")
var userService = data.NewUserService()

// Login functionality
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload request"})
		return
	}

	existingUser, err := userService.GetUserbyUsername(user.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Hash inputted password and check to see if it matches with the stored password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Wrong Password"})
		return
	}

	// Create a new token with claims including username and role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     existingUser.Role,
	})

	// Hash encoded token with the jwtSecret and append the result
	// to the token as a signature
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully logged in", "token": signedToken})

}

// Logged is a middleware that checks if the user is logged in
func Logged(c *gin.Context) {
	// Get the authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}
	// Split it into "Bearer" and <token>
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}
	// After validating the signing method, decode the header and payload that contains the claims
	// and re-hash them with the jwtSecret to ensure that the signature matches
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
	}

	c.Next()

}

// Admin is a middleware that checks if the user is an admin
// It checks the role of the user after decoding it from the token
// which is obtained from the authorization header
func Admin(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
	}

	// Check whether the role of the user is "admin"
	if token.Claims.(jwt.MapClaims)["role"] != "admin" {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		c.Abort()
	}

	c.Next()
}
