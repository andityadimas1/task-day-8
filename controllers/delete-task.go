package controllers

import (
		"time"
		"github.com/gin-gonic/gin"
		"fmt"
)

	type DeleteTask struct{
		Taskname    string 	  `json:"task name"`
		Deleted  	time.Time `json:"Deleted"`
	}

	func Deleted(c *gin.Context) {
		var del DeleteTask   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
		err := c.Bind(&del) // baut variable yg fungsinya untuk check error
		if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 72)
			fmt.Println("Terjadi error")
	}

		c.JSON(200, gin.H{
			"message"  : "Deleted",
			"data": map[string]interface{}{
				"Task Name": del.Taskname,
				"Deleted" : del.Deleted,
				},
			})
		}