package main

import (
	"log"
	"to-do-list/controllers"
	"to-do-list/config"
	"to-do-list/models"
	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connect()
	strDB := controllers.StrDB{DB: dbPG}
	request := gin.Default()

	models.Migrations(dbPG)
	// request.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//routes
	request.POST("/register", strDB.RegisterUser)
	request.POST("/login", strDB.LoginUser)
	request.GET("/getuser", strDB.GetDataUser)
	request.POST("/addtask", controllers.AddTask)
	// request.PUT("/updatetask", controllers.Update)
	log.Println("Server up and run on Port 8080")
	request.Run()
}