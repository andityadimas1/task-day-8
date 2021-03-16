package controllers

import (
	"net/http"
	"strconv"
	"time"
	"to-do-list/models"
	logger "to-do-list/sentry"

	"github.com/gin-gonic/gin"
)

func (StrDB *StrDB) LoginUser(c *gin.Context) {
	var (
		result gin.H
		user   models.User
	)

	Email, _ := strconv.ParseInt(c.PostForm("email"), 10, 64)
	Password, _ := strconv.ParseInt(c.PostForm("password"), 10, 64)
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
		logger.Sentry(err)

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
		user   User
	)
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		err := "not filled!"
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": err,
		})
		if res := StrDB.DB.Create(&user); res.Error != nil {
			err := res.Error
			result = gin.H{
				"status":  "Bad Request",
				"message": "Cant Process the Data!",
				"errors":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
			logger.Sentry(err)

		} else {
			StrDB.DB.Create(&user)
			result = gin.H{
				"status":  "success",
				"message": "Registered!",
				"data": map[string]interface{}{
					"id":       user.ID,
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
		logger.Sentry(err)
	} else {
		result = gin.H{
			"status": "success",
			"data":   user,
		}
		c.JSON(http.StatusOK, result)
	}
}

type User struct {
	ID          uint      `gorm:"primarykey, autoIncrement" json:"ID"`
	Email       string    `json:"email"`
	Password    string    `json:"passsword"`
	Name        string    `json:"nama"`
	Role        string    `json:"role"`
	CreatedDate time.Time `json:"id"`
}
