package mailer

import (
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"os"
)

func SendGoMail(to string, subject string, templateFile string) {
	// Lire le contenu du template HTML
	body, err := ioutil.ReadFile(templateFile)
	if err != nil {
		log.Fatalf("Erreur de lecture du template HTML : %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_USER"))
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(body))
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(
		os.Getenv("GMAIL_HOST"),
		587,
		os.Getenv("GMAIL_USER"),
		os.Getenv("GMAIL_PASSWORD"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
