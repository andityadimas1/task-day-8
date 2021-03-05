package controllers

	import (
		"time"
		"github.com/gin-gonic/gin"
		"fmt"
	)

	type UpdateTask struct{
		Taskname    string 	  `json:"task name"`
		Updated  	time.Time `json:"Updated"`
	}

	func Update(c *gin.Context) {
		var updt UpdateTask   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
		err := c.Bind(&updt) // baut variable yg fungsinya untuk check error
		if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 72)
			fmt.Println("Terjadi error")
	}

		c.JSON(200, gin.H{
			"message"  : "updated",
			"data": map[string]interface{}{
				"Task Name": updt.Taskname,
				"Deleted" : updt.Updated,
				},
			})
		}