package mail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)


const (
	authAddress   = "smtp.gmail.com"
	serverAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}


func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
) error {
	email := email.NewEmail()
	email.Subject = subject
	email.HTML = []byte(content)
	email.To = to
	email.Cc = cc
	email.Bcc = bcc

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword,authAddress)
	return email.Send(serverAddress, smtpAuth) 
}
