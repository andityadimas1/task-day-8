package models

import (
	"fmt"

	"gorm.io/gorm"
)

//function untuk migrasi ke database
func Migrations(db *gorm.DB) {
	if check := db.Migrator().HasTable(&User{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&User{})
		fmt.Println("Table berhasil tercreate")
	}
	if check := db.Migrator().HasTable(&Task{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&Task{})
		fmt.Println("Table berhasil tercreate")
	}
	if check := db.Migrator().HasTable(&ListData{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&ListData{})
		fmt.Println("Table berhasil tercreate")
	}
}
