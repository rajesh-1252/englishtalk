package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func UserRoute(rg *gin.RouterGroup) {
	rg.POST("/register", func(c *gin.Context) {
		var userReq UserRequest
		if err := c.ShouldBindJSON(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":    "user created successfully",
			"userDetail": userReq,
		})

	})
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": "OK TESTED",
		})
	})
}
