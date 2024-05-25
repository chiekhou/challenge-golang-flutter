package mailer

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"os"
	"text/template"
)

func SendGoMail(to string, subject string, templateFile string, data interface{}) {
	// Lire le contenu du template HTML
	var body bytes.Buffer
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
		return
	}

	err = t.Execute(&body, data)
	if err != nil {
		panic(err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer(
		os.Getenv("GMAIL_HOST"),
		587,
		os.Getenv("GMAIL_USER"),
		os.Getenv("GMAIL_PASSWORD"))

	// Envoyer l'email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
