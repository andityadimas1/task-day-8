package backgroundtask

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func RegisterEmail(email, message string) {

	to := []string{email}

	subject := "Succesfully"
	err := ConfigEmail(to, subject, message)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("terkirim!")
}

func ConfigEmail(to []string, subject, message string) error {

	err := godotenv.Load("env")
	if err == nil {
		body := "from" + os.Getenv("MAILER_NAME") + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Subject: " + subject + "\n\n" +
			message
		auth := smtp.PlainAuth("", os.Getenv("MAILER_NAME"), os.Getenv("MAILER_PASSWORD"), os.Getenv("MAILER_HOST"))
		smtpAdd := os.Getenv("MAILER_HOST") + ":" + os.Getenv("MAILER_PORT")
		err := smtp.SendMail(smtpAdd, auth, os.Getenv("MAILER_EMAIL"), to, []byte(body))

		if err == nil {
			return nil
		}
		return err
	}
	return err
}
