package routes

import (
	"englishTalk/controllers"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func UserRoute(rg *gin.RouterGroup) {
	rg.POST("/register", controllers.RegisterUser)
	rg.POST("/login", controllers.LoginUser)
}
