package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type RegisterRepository interface {
	CheckExistEmail(email string) (ok bool, err error)
	CreateSaveInfoRegister(info request.RegisterRequest) (infoRegister *model.SaveRegister, err error)
	CreateProfile(saveInfo model.SaveRegister) (profile *model.Profile, err error)
	CreateCredential(saveInfo *model.SaveRegister, profile *model.Profile, roleId uint) (err error)
	GetSaveInfo(saveInfoId uint) (saveInfo *model.SaveRegister, err error)
}
