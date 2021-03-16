package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func SeederAddtask(db *gorm.DB) {
	var AddtaskArray = [...][2]string{
		{"Makan Pagi", "true"},
		{"Mandi pagi", "true"},
		{"Pergi ke sekolah", "true"},
		{"Napping", "false"},
	}

	var task Task

	for _, data := range AddtaskArray {
		// Get Data from Array
		ID, _ := strconv.ParseInt(data[1], 10, 64)
		TaskNama, _ := strconv.ParseInt(data[2], 10, 64)
		task.Completed = data[0]
		db.Create(&task)
	}
	fmt.Println("Movie Data has been Seed!")
}
