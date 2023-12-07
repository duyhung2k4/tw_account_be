package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type RegisterRepository interface {
	CheckExistEmail(email string) (ok bool, err error)
	CreateSaveInfoRegister(info request.RegisterRequest) (infoRegister *model.SaveRegister, err error)
}
