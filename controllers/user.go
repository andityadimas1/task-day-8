package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"to-do-list/models"
	logger "to-do-list/sentry"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (StrDB *StrDB) LoginUser(c *gin.Context) {
	var (
		result gin.H
		user   models.User
	)

	Email := c.PostForm("email")
	Password := c.PostForm("password")
	// user := c.PostForm(user)

	// user.User = user
	user.Email = string(Email)
	user.Password = string(Password)

	if res := StrDB.DB.Create(&user); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status":  "Bad Request",
			"message": "Cant Process the Data!",
			"errors":  err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		logger.SentryString(err.Error())

	} else {
		Email := c.PostForm("email")
		Password := c.PostForm("password")

		StrDB.DB.Where(&user, "email = ? AND password = ?", Email, Password)

		result = gin.H{
			"status":  "success",
			"message": "Sucessfully Login!",
			// "data":    user,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (StrDB *StrDB) RegisterUser(c *gin.Context) {
	var (
		result gin.H
		user   models.User
	)
	if err := c.Bind(&user); err != nil || user.Email == "" || user.Password == "" || user.Name == "" || user.Role == "" {
		e := "Field Email, Password, FullName, Role is required!"
		result = gin.H{
			"status":  "bad request",
			"message": e,
		}
		fmt.Println("Field Email, Password, FullName, Role is required!")
		c.JSON(http.StatusBadRequest, result)

		logger.SentryString(err.Error())

	} else {
		if res := StrDB.DB.Create(&user); res.Error != nil {
			if check, data := GetRedis(key); check != false {
				if err := json.Unmarshal(data, &user); err != nil {
					fmt.Println("Error", err.Error())
				}
				fmt.Println(user)
				result = gin.H{
					"status":  "Bad Request",
					"message": "Cant Process the Data!",
					"errors":  err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)

				logger.SentryString(err.Error())

			} else {
				encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

				if err != nil {
					log.Println(err)
				}

				user.Password = string(encrypt)
				StrDB.DB.Create(&user)
				result = gin.H{
					"status":  "success",
					"message": "Registered!",
					"data": map[string]interface{}{
						// "id":       user.ID,
						"email":    user.Email,
						"fullName": user.Name,
						"role":     user.Role,
						"data":     user,
					},
				}
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
func (StrDB *StrDB) GetDataUser(c *gin.Context) {
	var (
		user   []models.User
		result gin.H
	)
	Email := c.Param("email")

	if res := StrDB.DB.Preload("email=", Email).Find(&user); res.Error != nil {
		err := res.Error
		result = gin.H{
			"status": "Not Found",
			"errors": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		logger.SentryString(err.Error())
	} else {
		result = gin.H{
			"status": "success",
			"data":   user,
		}
		c.JSON(http.StatusOK, result)
	}
}
