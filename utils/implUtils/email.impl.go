package impl_utils

import (
	"account-service/config"
	"account-service/utils"
	"net/smtp"
)

type emailUtils struct {
	emailFrom string
	smtpHost  string
	smtpPort  string
	password  string
}

func (e *emailUtils) SendEmailConfirmCodeRegister(email string, code string) (err error) {
	// Sender data.
	from := e.emailFrom
	password := e.password

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := e.smtpHost
	smtpPort := e.smtpPort

	// Message.
	message := []byte(code)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	errEmail := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	return errEmail
}

func EmailUtilsInit() utils.EmailUtils {
	return &emailUtils{
		emailFrom: config.GetEmailSend(),
		password:  config.GetPasswordEmailSend(),
		smtpPort:  config.GetSmtpPort(),
		smtpHost:  config.GetSmtpHost(),
	}
}
