package service

import (
	"account-service/dto/request"
	"account-service/model"
)

type LoginService interface {
	CheckCredential(info request.LoginRequest) (credential *model.Credential, err error)
	FindCredential(id uint) (credential *model.Credential, err error)
}
