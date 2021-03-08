package config

import(
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var userDB, passDB, hostDB, portDB, namaDB, ssl, timeZone string
	if err := godotenv.Load(".env"); err != nil{
		log.Fatal(err.Error())
	}else {
		userDB = os.Getenv("DATABASE")
		passDB = os.Getenv("DATABASE_PASS")
		hostDB = os.Getenv("DATABASE_HOST")
		portDB = os.Getenv("DATABSE_PORT")
		namaDB = os.Getenv("DATABASE_NAME")
		ssl = os.Getenv("DATABASE_SSLMODE")
		timeZone = os.Getenv("DATABASE_TIMEZONE")
	}

	conn :=
			" host=" + hostDB +
			" user=" + userDB +
			" password=" + passDB +
			" dbname=" + namaDB +
			" port=" + portDB +
			" sslmode=" + ssl +
			" TimeZone=" + timeZone

	// conn := "user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	
	// kondisi ketika akses database postgre
	if err != nil {
		//panic
		panic("Gagal masuk ke database")
	} else {
		fmt.Println("Koneksi ke database berhasil")
	}

	return db
}