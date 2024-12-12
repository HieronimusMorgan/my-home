package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	host := "smtp.example.com"
	port := "587"

	auth := smtp.PlainAuth("", from, password, host)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	return smtp.SendMail(fmt.Sprintf("%s:%s", host, port), auth, from, []string{to}, []byte(message))
}
