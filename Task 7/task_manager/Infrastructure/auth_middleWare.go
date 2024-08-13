package Infrastructure

import (
	"fmt"
	"strings"
	"task_manager/Domain"
	"task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var userservice = Usecases.NewUserService()

func Login(c *gin.Context) {
	var user Domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload request"})

		return
	}

	existingUser, err := userservice.GetUserbyUsername(user.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(existingUser.Username, existingUser.Role, existingUser.Password)
	if err := ComparePasswords(existingUser.Password, user.Password); err != nil {
		c.JSON(400, gin.H{"error": "Wrong Password"})
		return
	}

	signedToken, err := GenerateToken(user.Username, existingUser.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"message": "Successfully logged in", "token": signedToken})
}

func Logged(c *gin.Context) {

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

	token, err := ValidateToken(authParts[1])
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

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

	token, err := ValidateToken(authParts[1])
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if token.Claims.(jwt.MapClaims)["role"] != "admin" {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
