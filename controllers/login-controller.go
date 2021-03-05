package controllers

import (
		"fmt"
		"github.com/gin-gonic/gin"
)


	type LoginFunc struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	func Login(c *gin.Context) {
		var body LoginFunc
		err := c.Bind(&body)
		if err != nil {
			fmt.Println("Failed")
		}

		c.JSON(200, gin.H{
			"status"   	  :  "success",
			"access_token": "berhasil123",
		})
	}