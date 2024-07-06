package service

import (
	"fmt"
	"net/smtp"
)

const (
	from     = "blockchaindev100@gmail.com"
	smtpHost = "smtp.gmail.com"
	password = "xoeo rnfm xuzh adaa"
)

func SendMail(to string) {
	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte("To: " + to + "\r\n" +
		"Subject: User Verification\r\n" +
		"\r\n" +
		"Hello,\r\n" +
		"Click the Below link to verify the account\r\n")
	err := smtp.SendMail(smtpHost+":587", auth, from, []string{to}, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully")
}
