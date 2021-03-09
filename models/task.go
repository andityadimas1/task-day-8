package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uint   `json:"id"`
	TaskNama  string `json:"tasknama"`
	Completed string   `json:"completed"`
	// Created   time.Time `json:"created at"` 
}
