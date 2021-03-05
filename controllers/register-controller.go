package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

	type RegFunc struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nama 	 string `json:"nama"`
	}

	func Register(c *gin.Context) {
			var reg RegFunc   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
			err := c.Bind(&reg) // baut variable yg fungsinya untuk check error
			if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 72)
				fmt.Println("Terjadi error")
	}

<<<<<<< HEAD
	c.JSON(200, gin.H{
		"message"  : "added",
		"data": map[string]interface{}{
			"email":     reg.Email,
			"password":  reg.Password,
			"nama" : reg.Nama,
		},
	})
}
=======
		c.JSON(200, gin.H{
			"message"  : "added",
			"data": map[string]interface{}{
				"email":     reg.Email,
				"password":  reg.Password,
				"nama" : reg.Nama,
			},
		})
	}
>>>>>>> a6d7f90ad675674de398e7f1cdb8fe9836e8ba0f
