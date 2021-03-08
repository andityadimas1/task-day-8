package controllers

import (
	"fmt"
	"to-do-list/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) AddTask(c *gin.Context) {
	var (
		task  models.Task
		result gin.H
	)
	// fmt.Println(c.Bind(&users))
	if err := c.Bind(&task); err != nil {
		fmt.Println("NO Task Data")
	} else {
		StrDB.DB.Create(&task)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Added!",
			"data": map[string]interface{}{
				"id":       task.ID,
				"tasknama": task.TaskNama,
				"completed": task.Completed,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}
