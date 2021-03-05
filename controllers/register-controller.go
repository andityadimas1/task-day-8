package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

	type RegFunc struct {
		email    string `json:"email"`
		password string `json:"password"`
		nama 	 string `json:"nama"`
	}
	
	func Register(c *gin.Context) {
			var reg RegFunc   // buat variable untuk tipe data struct, dimana valuenya akan di isi di response
			err := c.Bind(&reg) // baut variable yg fungsinya untuk check error
			if err != nil {     // check error, if true -> masuk ke kondisi di line bawahnya (line 72)
				fmt.Println("Terjadi error")
	}

		c.JSON(200, gin.H{
			"message"  : "added",
			"data": map[string]interface{}{
				"email":     reg.email,
				"password":  reg.password,
				"nama" : reg.nama,
			},
		})
	}
