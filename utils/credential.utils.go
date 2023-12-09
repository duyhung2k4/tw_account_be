package utils

import "account-service/model"

type CredentialUtils interface {
	HashPassword(password string) (hashPassword string, err error)
	ComparePassword(password, hashPassword string) (success bool)
	GetRole() (roles []model.Role, err error)
}
