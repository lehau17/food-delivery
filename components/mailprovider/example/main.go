package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"text/template"
)

type EmailData struct {
	Email string
	Otp   string
}

func main() {

	from := "hau17131203@gmail.com"
	password := "quup slyz pzwp ifog"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	to := []string{"haudz20003@gmail.com"}

	subject := "Subject: Test Email from Golang\n"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	pw, _ := os.Getwd()
	path := filepath.Join(pw, "..", "..", "..", "storage", "templates", "email.html")
	emailBody, err := parseTemplate(path, EmailData{Email: to[0], Otp: "1234"})
	if err != nil {
		// return err
		log.Println(err)
	}
	// Gửi email
	message := []byte(subject + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + emailBody)
	// Gửi email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		// global.Logger.Error(err.Error())
		log.Println(err)
	}
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		// global.Logger.Error(err.Error())
		log.Println(err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		// global.Logger.Error(err.Error())
		log.Println(err)
		return "", err
	}

	return buf.String(), nil
}
