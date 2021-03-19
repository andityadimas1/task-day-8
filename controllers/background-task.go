package controllers

// func SendEmailConfiguration(to []string, cc []string, subject, message string) {

// 	err := godotenv.Load(".env")
// 	if err == nil {
// 		body := "From: " + os.Getenv("email_email") + "\n" +
// 			"To: " + strings.Join(to, ",") + "\n" +
// 			"Cc : " + strings.Join(cc, ",") + "\n" +
// 			"Subject: " + subject + "\n\n" +
// 			message

// 		auth := smtp.PlainAuth("", os.Getenv("email_email"), os.Getenv("email_password"), os.Getenv("email_smptp_hostname"))
// 		smtpAdd := os.Getenv("email_smptp_hostname") + ":" + os.Getenv("email_smptp_port")
// 		err := smtp.SendMail(smtpAdd, auth, os.Getenv("email_email"), append(to, cc...),
// 			[]byte(body))

// 		if err == nil {
// 			return nil
// 		}
// 		return err
// 	}
// 	return err
// }

// func SendEmail() {
// 	to := []string{"sahlan.nasution07@gmail.com", "andityadimas@gmail.com"}
// 	cc := []string{"dimas.aninditya@xapiens.id"}

// 	subject := "Test email notification"
// 	message := "Halo gaes, ini tes email notif ya"

// 	err := SendEmailConfiguration(to, cc, subject, message)
// 	if err != nil { // kalau ada error saat send email
// 		log.Println(err.Error())
// 	}

// 	log.Println("Mail send!")
// }
