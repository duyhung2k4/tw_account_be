package config

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	appPort = os.Getenv(APP_PORT)
	dbName = os.Getenv(DB_NAME)
	dbUser = os.Getenv(DB_USER)
	dbPort = os.Getenv(DB_PORT)
	dbHost = os.Getenv(DB_HOST)
	dbPassword = os.Getenv(DB_PASSWORD)

	smtpHost = os.Getenv(SMTP_HOST)
	smtpPort = os.Getenv(SMTP_PORT)
	emailSend = os.Getenv(EMAIL_SEND)
	passwordEmailSend = os.Getenv(PASSWORD_EMAIL_SEND)

	return nil
}
