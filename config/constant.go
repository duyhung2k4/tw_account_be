package config

import (
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

var (
	appPort    string
	dbName     string
	dbUser     string
	dbPort     string
	dbHost     string
	dbPassword string
	db         *gorm.DB

	smtpHost          string
	smtpPort          string
	emailSend         string
	passwordEmailSend string

	tokenAuth *jwtauth.JWTAuth
)

const (
	APP_PORT    = "APP_PORT"
	DB_NAME     = "DB_NAME"
	DB_USER     = "DB_USER"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_PASSWORD = "DB_PASSWORD"

	SMTP_HOST           = "SMTP_HOST"
	SMTP_PORT           = "SMTP_PORT"
	EMAIL_SEND          = "EMAIL"
	PASSWORD_EMAIL_SEND = "EMAIL_PASSWORD"
)
