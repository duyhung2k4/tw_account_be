package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	"account-service/model"
	"account-service/repository"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"

	"gorm.io/gorm"
)

type registerRepository struct {
	db              *gorm.DB
	registerUtils   utils.RegisterUtils
	credentialUtils utils.CredentialUtils
}

func (r *registerRepository) CheckExistEmail(email string) (ok bool, err error) {
	var creadential *model.Credential

	errCredential := r.db.Model(&model.Credential{}).Where("email = ?", email).Find(&creadential).Error

	if errCredential != nil {
		return false, errCredential
	}

	if creadential != nil {
		return true, nil
	}

	return false, nil
}

func (r *registerRepository) CreateSaveInfoRegister(info request.RegisterRequest) (infoRegister *model.SaveRegister, err error) {
	password, errHashPassword := r.credentialUtils.HashPassword(info.Password)
	code := r.registerUtils.CreateCode()

	if errHashPassword != nil {
		return nil, errHashPassword
	}

	startAt, finishAt := r.registerUtils.CreateTimeExist()

	saveRegister := &model.SaveRegister{
		Code:     code,
		Username: info.Username,
		Email:    info.Email,
		Password: password,
		Role:     info.Role,
		StartAt:  startAt,
		FinishAt: finishAt,
	}

	errCreateSaveRegister := r.db.Model(&model.SaveRegister{}).Create(&saveRegister).Error

	if errCreateSaveRegister != nil {
		return nil, errCreateSaveRegister
	}

	return saveRegister, nil
}

func RegisterRepositoryInit() repository.RegisterRepository {
	db := config.GetDB()
	registerUtils := impl_utils.RegisterUtilsInit()
	credentialUtils := impl_utils.CredentialUtilsInit()
	return &registerRepository{
		db:              db,
		registerUtils:   registerUtils,
		credentialUtils: credentialUtils,
	}
}
