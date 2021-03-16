package models

import (
	"fmt"

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
		task.ID = 0
		task.TaskNama = data[0]
		task.Completed = data[0]
		db.Create(&task)
	}
	fmt.Println("seed!")
}

func SeederUser(db *gorm.DB) {
	var userArray = [...][4]string{
		{"DImmas.anin@xapiens.id", "saya123", "Dimas ganteng", "admin"},
		{"sahlannasution@gmail.com", "inisahlan", "Sahlan ganteng", "guest"},
		{"dimasdimas@gmail.com", "dimdim1234", "dimas", "guest"},
	}

	func SeederUser(db *gorm.DB) {
		var userArray = [...]string{
			"admin",
			"test",
			"guest",
		}
	
		var user Users
	
		for _, v := range userArray {
			user.Username = v
			user.Role = v
			user.Password = v
	
			// enkrip password
			// param 1 dari password yang sudah ditentukan (contoh : admin)
			// kita bikin konfersi dari string ke byte -> caranya []byte(user.Password)
			encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	
			// checking apakah proses enkripsi error / tidak
			if err != nil {
				log.Println(err)
			}
	
			user.Password = string(encrypt)
			user.ID = 0 // declare id dimulai dari 0, karena auto increment
			db.Create(&user)
		}
		fmt.Println("Seeder user created")
	}