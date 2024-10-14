package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
)

type EmailConfig struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func NewEmailConfig(from, password, smtpHost, smtpPort string) *EmailConfig {
	return &EmailConfig{from: from, password: password, smtpHost: smtpHost, smtpPort: smtpPort}
}

func (e *EmailConfig) ParseTemplate(templateFileName string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (e *EmailConfig) SendMail(to []string, templateName, subject string, data interface{}) error {
	// from := "hau17131203@gmail.com"
	// password := "quup slyz pzwp ifog"
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"

	auth := smtp.PlainAuth("", e.from, e.password, e.smtpHost)
	pw, _ := os.Getwd()
	path := filepath.Join(pw, "storage", "templates", templateName)
	// EmailData{Email: to[0], Otp: common.GetOtp()
	emailBody, err := e.ParseTemplate(path, data)
	if err != nil {
		return err
	}
	// Gửi email
	message := []byte(subject + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + emailBody)
	// Gửi email
	err = smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
