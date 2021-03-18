package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateTask struct {
	Taskname string    `json:"task name"`
	Updated  time.Time `json:"Updated"`
}

func Update(c *gin.Context) {
	var updt UpdateTask
	err := c.Bind(&updt)
	if err != nil {
		fmt.Println("Terjadi error")
	}

	c.JSON(200, gin.H{
		"message": "updated",
		"data": map[string]interface{}{
			"Task Name": updt.Taskname,
			"Deleted":   updt.Updated,
		},
	})
}
