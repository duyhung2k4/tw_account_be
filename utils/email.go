package utils

type EmailUtils interface {
	SendEmailConfirmCodeRegister(email string, code string) (err error)
}
