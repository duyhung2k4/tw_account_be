package config

import (
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

func GetAppPort() string {
	return appPort
}

func GetDB() *gorm.DB {
	return db
}

func GetJWT() *jwtauth.JWTAuth {
	return tokenAuth
}

func GetSmtpHost() string {
	return smtpHost
}

func GetSmtpPort() string {
	return smtpPort
}

func GetEmailSend() string {
	return emailSend
}

func GetPasswordEmailSend() string {
	return passwordEmailSend
}
