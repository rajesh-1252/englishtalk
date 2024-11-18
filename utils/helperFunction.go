package utils

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ApiResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    data,
	})
}

func GetJWT(c *gin.Context) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is not set"})
		return "", fmt.Errorf("JWT secret is not set")
	}
	return jwtSecret, nil
}

func GenerateJWT(userID uint, secret string) (string, error) {
	// Define token claims
	expires, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		expires = 1
	}
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * time.Duration(expires)).Unix(), // Token expires in 24 hours
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString([]byte(secret))
}
