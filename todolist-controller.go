package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	
	request := gin.Default()

	//test postman
	request.GET("/task", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	request.GET("/task/login/", func(c *gin.Context) {
		email := c.Param("email") // ini code yg membedakan query & parameter
		password := c.Param("password")
		c.JSON(200, gin.H{
			"email": email,
			"password":password,
		})
	})

	request.POST("task/register/", func (c *gin.Context)  {
		nama := c.Param("nama")
		email := c.Param("email")
		password := c.Param("password")
		c.JSON(200, gin.H{
			"nama" : nama,
			"email": email,
			"password":password,
		})
	})	
	
	request.POST("/todolist/", func(c *gin.Context) {
		email := c.PostForm("email")
		taskname := c.PostForm("task name")
		completed := c.PostForm("completed")
		time := c.PostForm("time")

		c.JSON(200, gin.H{
			"email": email,
			"task nname":  taskname,
			"completed": completed,
			"time": time,
		})
	})

}

