package controllers

import(
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

	type AddFunc struct {
		Taskname    string 	  `json:"task name"`
		Completed	bool   	  `json:"completed"`
		Created	 	time.Time `json:"created"`
	}

	func AddTask (c *gin.Context) {
		var add AddFunc
		err := c.Bind(&add) 
		if err != nil { 
			fmt.Println("Terjadi error")
	}

	c.JSON(200, gin.H{
		"message"  : "task added",
		"data": map[string]interface{}{
			"Task Name":  add.Taskname,
			"Completed": add.Completed,
			"Created" : add.Created,
		},
	})
}	