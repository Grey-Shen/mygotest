package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "1518522971@qq.com", "xiaoyu521JUN", "mail.qq.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"2593929657@qq.com"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.qq.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
