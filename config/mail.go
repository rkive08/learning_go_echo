// package config

// import (
// 	"bytes"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/smtp"
// 	"os"
// )

// type Mailer struct {
// 	Host string
// 	Port string
// 	User string
// 	Pass string
// 	From string
// }

// var Mail Mailer

// func LoadMailConfig() {
// 	Mail = Mailer{
// 		Host: os.Getenv("SMTP_HOST"),
// 		Port: os.Getenv("SMTP_PORT"),
// 		User: os.Getenv("SMTP_USER"),
// 		Pass: os.Getenv("SMTP_PASS"),
// 		From: os.Getenv("SMTP_FROM"),
// 	}
// }

// // for test
// func SendEmail(to string, subject string, body string) error {
// 	auth := smtp.PlainAuth("", Mail.User, Mail.Pass, Mail.Host)

// 	msg := []byte(
// 		"From: " + Mail.From + "\r\n" +
// 			"To: " + to + "\r\n" +
// 			"Subject: " + subject + "\r\n\r\n" +
// 			body,
// 	)

// 	addr := fmt.Sprintf("%s:%s", Mail.Host, Mail.Port)

// 	err := smtp.SendMail(addr, auth, Mail.From, []string{to}, msg)
// 	if err != nil {
// 		log.Println("Gagal kirim email:", err)
// 		return err
// 	}
// 	log.Println("Email terkirim ke:", to)
// 	return nil
// }

// // end test

// func SendEmailTemplate(to string, subject string, templateFile string, data interface{}) error {
// 	// parse template
// 	tmpl, err := template.ParseFiles(templateFile)
// 	if err != nil {
// 		return err
// 	}

// 	var body bytes.Buffer

// 	// header
// 	body.Write([]byte("MIME-Version: 1.0\r\n"))
// 	body.Write([]byte("Content-Type: text/html; charset=\"UTF-8\"\r\n"))
// 	body.Write([]byte(fmt.Sprintf("From: %s\r\n", Mail.From)))
// 	body.Write([]byte(fmt.Sprintf("To: %s\r\n", to)))
// 	body.Write([]byte(fmt.Sprintf("Subject: %s\r\n\r\n", subject)))

// 	// isi template
// 	err = tmpl.Execute(&body, data)
// 	if err != nil {
// 		return err
// 	}

// 	// auth
// 	auth := smtp.PlainAuth("", Mail.User, Mail.Pass, Mail.Host)
// 	addr := fmt.Sprintf("%s:%s", Mail.Host, Mail.Port)

// 	// kirim email
// 	err = smtp.SendMail(addr, auth, Mail.From, []string{to}, body.Bytes())
// 	if err != nil {
// 		log.Println("Gagal kirim email:", err)
// 		return err
// 	}
// 	log.Println("Email terkirim ke:", to)
// 	return nil
// }

package config

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	gomail "github.com/go-mail/mail/v2"
)

type Mailer struct {
	Host string
	Port int
	User string
	Pass string
	From string
}

var Mail Mailer

func LoadMailConfig() {
	// convert port string ke int
	port := 587
	if os.Getenv("SMTP_PORT") != "" {
		fmt.Sscanf(os.Getenv("SMTP_PORT"), "%d", &port)
	}

	Mail = Mailer{
		Host: os.Getenv("SMTP_HOST"),
		Port: port,
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
		From: os.Getenv("SMTP_FROM"),
	}
}

func SendEmailTemplate(to string, subject string, templateFile string, data interface{}) error {
	// parse template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return err
	}

	// buat pesan
	m := gomail.NewMessage()
	m.SetHeader("From", Mail.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	// dialer
	d := gomail.NewDialer(Mail.Host, Mail.Port, Mail.User, Mail.Pass)

	// kirim
	if err := d.DialAndSend(m); err != nil {
		log.Println("Gagal kirim email:", err)
		return err
	}

	log.Println("Email terkirim ke:", to)
	return nil
}
