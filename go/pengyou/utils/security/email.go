package security

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(
	host string,
	port int,
	username string,
	password string,
	sender string,
	to string,
	subject string,
	body string) error {

	m := gomail.NewMessage()

	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, password)

	err := d.DialAndSend(m)

	return err
}
