package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"Clifton",
		"cliftonavil@gmail.com",
		"*********",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"cliftonavil@gmail.com",
		[]string{"cliftonavil@gmail.com"},
		[]byte("To: recipient@example.net\r\n"+
			"Subject: discount Gophers!\r\n"+
			"\r\n"+
			"This is the email body.\r\n"),
	)
	if err != nil {
		log.Fatal(err)
	}
}
