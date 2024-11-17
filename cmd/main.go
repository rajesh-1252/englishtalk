package main

import (
	"englishTalk/routes"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

func main() {
	var age int = 1
	age = 2
	println(age)
	r := gin.Default()
	apiV1 := r.Group("api/v1")
	{
		routes.UserRoute(apiV1)
	}
	r.Run(":8081")
}
