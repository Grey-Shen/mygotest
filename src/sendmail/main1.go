package main

import (
	"log"
	"net/smtp"
)

// var (
// 	subject  = flag.String("s", "", "subject of the mail")
// 	body     = flag.String("b", "", "body of themail")
// 	reciMail = flag.String("m", "", "recipient mail address")
// )

// func main() {
// 	// Set up authentication information.
// 	flag.Parse()
// 	sub := fmt.Sprintf("subject: %s\r\n\r\n", *subject)
// 	content := *body
// 	mailList := strings.Split(*reciMail, ",")

// 	auth := smtp.PlainAuth(
// 		"",
// 		"1518522971@qq.com",
// 		"logudhmiyzqtiffc",
// 		"smtp.qq.com",
// 		//"smtp.gmail.com",
// 	)
// 	// Connect to the server, authenticate, set the sender and recipient,
// 	// and send the email all in one step.
// 	err := smtp.SendMail(
// 		"smtp.qq.com:25",
// 		auth,
// 		"1518522971@qq.com",
// 		mailList,
// 		[]byte(sub+content),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func main() {
	hostname := "smtp.	qq.com"
	to := []string{"2593929657@qq.com"}
	from := "1518522971@qq.com"
	passwd := "logudhmiyzqtiffc"
	auth := smtp.PlainAuth("", from, passwd, hostname)

	err := smtp.SendMail(hostname+":465", auth, from, to, []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
}
