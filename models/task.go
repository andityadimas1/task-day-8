package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Id        uint   `json:"id"`
	TaskNama  string `json:"tasknama"`
	Completed bool `json:"completed"`
	// Created   time.Time `json:"created at"` 
}
