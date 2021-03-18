package models

import (
	"gorm.io/gorm"
)

type ListData struct {
	gorm.Model
	ID        uint   `json:"id"`
	ListNama  string `json:"listnama"`
	Completed string `json:"completed"`
	// IdList uint `gorm :"foreignKey:"id" json: Idlist`
	// Created   time.Time `json:"created at"`
}
