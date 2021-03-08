package controllers

import "gorm.io/gorm"

type StrDB struct {
	DB *gorm.DB
}