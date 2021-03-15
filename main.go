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
	request.GET("/getuser/:ID", strDB.GetDataUser)
	request.POST("/addtask", strDB.AddTask)
	request.POST("/updatetask/:ID", strDB.UpdateTask)
	request.DELETE("/deletetask/:ID", strDB.DeleteTask)
	request.PUT("/gettask/:ID", strDB.GetTask)
	// request.PUT("/listtask/", strDB.GetTask)
	log.Println("Server up and run on Port 8080")
	request.Run()
}