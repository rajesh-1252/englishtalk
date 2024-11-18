package main

import (
	"englishTalk/database"
	"englishTalk/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		println("Error loading .env file: %v", err)
	}
	database.ConnectDb()
	r := gin.Default()
	apiV1 := r.Group("api/v1")
	{
		routes.UserRoute(apiV1)
	}
	r.Run(":8081")
}
