package main

import (
	"log"
	"to-do-list/config"
	"to-do-list/controllers"
	"to-do-list/models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connect()
	strDB := controllers.StrDB{DB: dbPG}

	// models to db
	models.Migrations(dbPG)

	//seeding data
	models.SeederUser(dbPG)
	models.SeederAddtask(dbPG)

	request := gin.Default()

	request.POST("/register", strDB.RegisterUser)

	request.POST("/login", strDB.MiddleWare().LoginHandler)

	request.NoRoute(strDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := request.Group("/auth")

	auth.Use(strDB.MiddleWare().MiddlewareFunc())
	{
		auth.POST("/addtask", strDB.AddTask)
		auth.POST("/updatetask/:ID", strDB.UpdateTask)
		auth.DELETE("/deletetask/:ID", strDB.DeleteTask)
		auth.PUT("/gettask/:ID", strDB.GetTask)
		auth.PUT("/getuser", strDB.GetDataUser)
	}

	// request.POST("/register", strDB.RegisterUser)
	// request.POST("/login", strDB.LoginUser)

	// request.GET("/getuser/:ID", strDB.GetDataUser)
	// request.POST("/addtask", strDB.AddTask)
	// request.POST("/updatetask/:ID", strDB.UpdateTask)
	// request.DELETE("/deletetask/:ID", strDB.DeleteTask)
	// request.PUT("/gettask/:ID", strDB.GetTask)
	// request.PUT("/listtask/", strDB.GetTask)

	log.Println("Server up and run on Port 8080")
	request.Run()
}
