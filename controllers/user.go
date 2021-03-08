package controllers

import (
	
	"time"
	"fmt"
	"to-do-list/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) LoginUser(c *gin.Context) {
	var (
		// Id    models.User
		result gin.H
		user models.User
	)
	if  err := c.Bind(&user); err != nil{
		fmt.Println("Tidak dapat login")
	} else {
		Email := c.PostForm("email")
		Password := c.PostForm("password")

		StrDB.DB.Where(&user, "email = ? AND password = ?", Email, Password)
		
		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Login!",
			"data" : user,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (StrDB *StrDB) RegisterUser(c *gin.Context) {
	var (
		result gin.H
		user models.User
	)
		if  err := c.Bind(&user); err != nil{
		fmt.Println("Tidak dapat resgistrasi")
	} else {
		StrDB.DB.Create(&user)
		result = gin.H{
			"status":  "success",
			"message": "Registered!",
			"data": map[string]interface{}{
				"id":       user.ID,
				"email":    user.Email,
				"fullName": user.Name,
				"data" : user,
			},
		}
	}
	c.JSON(http.StatusOK, result)
}

func (StrDB *StrDB) GetDataUser(c *gin.Context){
		var (
			result gin.H
			user models.User

		)
		name := c.Param("name = ?")

		StrDB.DB.Find(&user, "name = ?", name)
		
		result = gin.H{
			"status":  "success",
			"message": "Catch IT!",
			// "data" : user,
		}
		c.JSON(http.StatusOK, result)		
}
type User struct{
		Id string `json:"ID"`
		Email string `json:"email"`
		Password string `json:"passsword"`
		Name string `json:"nama"`
		CreatedDate time.Time `json:"id"`
}


// func (strDB *StrDB) Register(*gin.Context){
// 	var (
// 		Email  models.User
// 		Password models.User
// 		User models.User
// 	)
	
