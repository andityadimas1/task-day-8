package models

import (
	"fmt"

	"gorm.io/gorm"
)

//function untuk migrasi ke database
func Migrations(db *gorm.DB) {


	if check := db.Migrator().HasTable(&User{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&User{})
		fmt.Println("Table vendor berhasil di create")
	}
}