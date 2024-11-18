package controllers

import (
	"net/http"
	"os"

	"englishTalk/database"
	"englishTalk/error"
	"englishTalk/models"
	"englishTalk/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var req models.User

	// Bind JSON input to the user struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Get the JWT secret from environment variables
	jwtSecret := os.Getenv("JWT_SECRET")
	println(jwtSecret, "")
	if jwtSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is not set"})
		return
	}

	// Combine the JWT secret with the password before hashing
	passwordWithSecret := req.Password + jwtSecret

	// Hash the combined password and secret
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSecret), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	req.Password = string(hashedPassword)

	// Save the user to the database
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	// Respond with success message and user details (excluding password)
	req.Password = "" // Omit password in response
	c.JSON(http.StatusOK, gin.H{
		"message":    "User created successfully",
		"userDetail": req,
	})
}

func LoginUser(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req Login
	var user models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		error.BadRequestError(c, "Invalid input", err.Error())
		return
	}

	// Fetch user from the database
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		error.BadRequestError(c, "Invalid email or password", err.Error())
		return
	}

	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		error.BadRequestError(c, "JWT secret is not set", "JWT_SECRET is missing in environment")
		return
	}

	// Validate the password
	passwordWithSecret := req.Password + jwtSecret
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordWithSecret)); err != nil {
		error.BadRequestError(c, "Invalid email or password", err.Error())
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateJWT(user.Id, jwtSecret)
	if err != nil {
		error.BadRequestError(c, "Failed to generate token", err.Error())
		return
	}

	// Respond with the token
	user.Password = ""
	utils.ApiResponse(c, "Login successful", gin.H{"user": user, "token": token})
}

func UpdateUserDetail(c *gin.Context) {

}

func UpdatePassword(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
