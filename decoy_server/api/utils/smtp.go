package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(message string) {
	password := GoDotEnvVariable("PASSWD")
	from := GoDotEnvVariable("EMAIL")
	toList := []string{"ankithans1947@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"

	fro := fmt.Sprintf("From: <%s>\r\n", "dserver03@gmail.com")
	to := fmt.Sprintf("To: <%s>\r\n", "ankithans1947@gmail.com")
	subject := "Subject: no-reply-decoy-server-team\r\n"

	msg := fro + to + subject + "\r\n" + message
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(ColorBlue), "Successfully sent mail to the responsible team", string(ColorReset))
}
