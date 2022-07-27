package controller

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func sendMail(mailDetails map[string]string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "paulsimonk2@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", mailDetails["to"])

	// Set E-Mail subject
	m.SetHeader("Subject", mailDetails["subject"])

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", mailDetails["body"])

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "paulsimonk2@gmail.com", "saqyvaklpauefeoc")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
