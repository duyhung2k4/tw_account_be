package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type LoginRepository interface {
	FindCredential(info request.LoginRequest) (credential *model.Credential, err error)
}
