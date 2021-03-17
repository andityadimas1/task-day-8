package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       uint   `json:"id"`
	Email    string `gorm:"primarykey" json:"email"`
	Password string `json:"password"`
	Name     string `json:"fullName"`
	Role     string `json:"role"`
}
