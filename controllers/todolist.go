package controllers

import (
	"fmt"
	"net/http"
	"to-do-list/models"
	logger "to-do-list/sentry"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) AddTask(c *gin.Context) {
	var (
		task   models.Task
		result gin.H
	)
	// tasknama := c.PostForm("tasknama")
	// completed := c.PostForm("completed")

	// task.TaskNama = tasknama
	// task.Completed = completed

	if err := c.Bind(&task); err != nil || task.TaskNama == "" || task.Completed == "" {
		e := "ADA YANG BELUM DIISI!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		fmt.Println("Field ada yang belum diisi")
		c.JSON(http.StatusBadRequest, result)

		logger.Sentry(err) // push log error ke sentry

	} else {

		if res := StrDB.DB.Create(&task); res.Error != nil {
			err := res.Error
			result = gin.H{
				"status":  "Bad Request",
				"message": "Cant Process the Data!",
				"errors":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)

			logger.Sentry(err)
		} else {
			StrDB.DB.Create(&task)
			result = gin.H{
				"status":  "success",
				"message": "Sucessfully Added!",
				"data": map[string]interface{}{
					// "ID":        task.ID,
					"tasknama":  task.TaskNama,
					"completed": task.Completed,
					"data":      task,
				},
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

func (StrDB *StrDB) UpdateTask(c *gin.Context) {
	var (
		task   models.Task
		result gin.H
	)

	tasknama := c.Param("TaskNama")
	completed := c.Param("Completed")

	task.TaskNama = tasknama
	task.Completed = completed

	if res := StrDB.DB.Where("tasknama = ?", tasknama).First(&task); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "not found",
			"message": "Task not found!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err)
	} else {
		task.TaskNama = tasknama
		StrDB.DB.Save(&task)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Updated Data!",
			"data":    task,
		}
		c.JSON(http.StatusOK, result)
	}
}

func (strDB *StrDB) DeleteTask(c *gin.Context) {
	var (
		task []models.Task
	)
	ID := c.Param("ID")
	d := strDB.DB.Where("id", ID).Delete(&task)
	fmt.Println(d)
	c.JSON(200, gin.H{"ID" + ID: "deleted"})
}

func (StrDB *StrDB) GetTask(c *gin.Context) {
	var (
		task   []models.Task
		result gin.H
	)

	if res := StrDB.DB.Preload("Task").Find(&task); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Not Found",
			"message": "Cannot Get Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Get Data!",
			"data":    task,
		}
		c.JSON(http.StatusOK, result)
	}

	// func (StrDB *StrDB) GetAllListTask(c *gin.Context) {
	// 	var (
	// 		task []models.Task
	// 		result gin.H
	// 	)
	// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// 	paginator := helpers.Paging (&helpers.Param{
	// 		DB:      strDB.DB,
	// 		Page:    page,
	// 		Limit:   limit,
	// 		OrderBy: []string{"id task"},
	// 		ShowSQL: true,
	// 		Join:    "",
	// 		Query:   "",
	// 	}
	// StrDB.DB.Find(&task)
	// result = gin.H{
	// 	"status":  "success",
	// 	"message": "Successfully Listed",
	// 	"data":    task,
	// }
	// 	c.JSON(http.StatusOK, result)
	//  }

}
