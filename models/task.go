package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uint   `json:"id"`
	TaskNama  string `json:"tasknama"`
	Completed string `json:"completed"`
	// IdList uint `gorm :"foreignKey:"id" json: Idlist`
	// Created   time.Time `json:"created at"` 
}
