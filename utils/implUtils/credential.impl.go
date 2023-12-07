package impl_utils

import (
	"account-service/utils"

	"golang.org/x/crypto/bcrypt"
)

type credentialUtils struct{}

func (c *credentialUtils) HashPassword(password string) (hashPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (c *credentialUtils) ComparePassword(password, passwordHash string) (success bool) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func CredentialUtilsInit() utils.CredentialUtils {
	return &credentialUtils{}
}
