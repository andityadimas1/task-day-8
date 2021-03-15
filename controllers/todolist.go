package controllers

import (
	"fmt"
	"to-do-list/models"
	"net/http"
	"github.com/gin-gonic/gin"
	// "strconv"
)

func (StrDB *StrDB) AddTask(c *gin.Context) {
	var (
		task  models.Task
		result gin.H
	)
	if err := c.Bind(&task); err != nil {
		fmt.Println(err.Error())
	} else {
		StrDB.DB.Create(&task)
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Added!",
			"data": map[string]interface{}{
				"ID":       task.ID,
				"tasknama": task.TaskNama,
				"completed": task.Completed,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (strDB *StrDB) UpdateTask(c *gin.Context) {
	var (
		task  models.Task
		result gin.H
	)
		ID := c.Param("ID")
		TaskNama := c.Param("TaskNama")
		Completed := c.Param("Completed")

	if err := c.Bind(&task); err != nil {
		fmt.Println("No Data or something wrong happen!!!")
	} else {
		// 	d := strDB.DB.Where("id", ID).Delete(&task)
		strDB.DB.Where("id", ID).Find(&task)
		task.TaskNama = TaskNama
		task.Completed = Completed

		result = gin.H{
			"message": "success Update Data",
			"data":    task,
		}
		strDB.DB.Save(&task)
		c.JSON(http.StatusOK, result)
	}
}

// func (StrDB *StrDB) UpdateTask(c *gin.Context) {
// 	var(
// 		task  models.Task
// 		result gin.H
// 	)

// 	TaskNama := c.Param("TaskNama")
// 	ID := c.Param("ID")

// 	StrDB.DB.Where("ID = ?", ID).Find(&task)
// 	task.TaskNama = TaskNama
// 	StrDB.DB.Save(&task)
// 	result = gin.H{
// 		"status":  "success",
// 		"message": "Task Sucessfully updated!",
// 		"data":    task,
// 	}

// 	c.JSON(http.StatusOK, result)
// }

func (strDB *StrDB) DeleteTask(c *gin.Context) {
	var(
		task [] models.Task
	)
	ID := c.Param("ID")
	d := strDB.DB.Where("id", ID).Delete(&task)
	fmt.Println(d)
	c.JSON(200, gin.H{"ID" + ID: "deleted"})
}


func (StrDB *StrDB) GetTask(c *gin.Context) {
	var (
		task [] models.Task
		result gin.H
	)
	// ID := c.Query("ID")
	// StrDB.DB.First(&task, ID)
	StrDB.DB.Preload("ID").Find(&task)
	result = gin.H{
		"status":  "success",
		"message": "Successfully",
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

type Task struct{
	ID  string `gorm:"primarykey" json:"id"`
	Tasknama string `json:"tasknama"`
	Completed string   `json:"completed"`
}