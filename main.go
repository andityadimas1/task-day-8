package main

import (
	"log"
	"to-do-list/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	request := gin.Default()

	request.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	request.POST("/register", controllers.Register)
	request.POST("/login", controllers.Login)
	log.Println("Server up and run on Port 8080")
	request.Run()
}