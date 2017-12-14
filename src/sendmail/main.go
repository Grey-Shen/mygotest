package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1518522971@qq.com")
	m.SetHeader("To", "2593929657@qq.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer("smtp.qq.com", 465, "1518522971@qq.com", "xiaoyu521JUN")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("send failed err: ", err)
	}
}
