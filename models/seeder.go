package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederAddtask(db *gorm.DB) {
	var AddtaskArray = [...][2]string{
		{"Makan Pagi", "true"},
		{"Mandi pagi", "true"},
		{"Pergi ke sekolah", "true"},
		{"Napping", "false"},
		{"main kelereng", "false"},
		{"Nonton netflix", "false"},
		{"main dota2 sampai pagi", "true"},
	}

	var task Task

	for _, data := range AddtaskArray {
		// Get Data from Array
		task.ID = 0
		task.TaskNama = data[0]
		task.Completed = data[1]
		db.Create(&task)
	}
	fmt.Println("seed!")
}

func SeederUser(db *gorm.DB) {
	var userArray = [...][4]string{
		{"Dimas.anin@xapiens.id", "saya123", "Dimas ganteng", "admin"},
		{"dimdimdimidmidmidmdimdimdim@gmail.com", "dimas1234455", "dimas1111", "guest"},
		{"dimasdimas@gmail.com", "dimdim1234", "dimas", "guest"},
		{"dimaslikescoffee@gmail.com", "dimas likes coffee", "dimas", "guest"},
		{"dimasalwayslikescoffee@gmail.com", "dimas always likes coffee", "dimaskok", "guest"},
		{"dimaslike@gmail.com", "dimas coffee", "dimas ya", "guest"},
		{"dimastataga@gmail.com", "dimas tatagca", "asaas", "guest"},
	}

	var user User

	for _, v := range userArray {
		// user.ID = 0
		user.Role = v[3]
		user.Password = v[1]
		user.Email = v[0]
		user.Name = v[2]

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
