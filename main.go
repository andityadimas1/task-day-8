package main

import (
	"fmt"
	"log"
	"time"
	"to-do-list/backgroundtask"
	"to-do-list/config"
	"to-do-list/controllers"
	"to-do-list/models"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
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
		auth.GET("/getuser", strDB.GetDataUser)
	}
	log.Println("Server up and run on Port 8080")
	// request.Run()

	go func() {
		gocron.Every(30).Minutes().Do(Crownjob)
		<-gocron.Start()
	}()

	request.Run(":8080")
}

func Crownjob() {
	var (
		mail []models.RegistEmail
	)
	dbPG := config.Connect()
	strDB := controllers.StrDB{DB: dbPG}

	strDB.DB.Where("status = ?", false).Find(&mail)

	fmt.Println(len(mail))

	for i := 0; i < len(mail); i++ {
		backgroundtask.RegisterEmail(mail[i].Email, mail[i].Message)
		mail[i].Status = true
		mail[i].DeliveredAt = time.Now()
		strDB.DB.Save(&mail[i])
	}
}
