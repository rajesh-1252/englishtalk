package routes

import (
	"englishTalk/database"
	"englishTalk/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func UserRoute(rg *gin.RouterGroup) {
	rg.POST("/register", func(c *gin.Context) {
		var req models.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		req.Password = string(hashedPassword)

		result := database.DB.Create(&req)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message":    "user created successfully",
			"userDetail": req,
		})
	})

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": "OK TESTED",
		})
	})
}
