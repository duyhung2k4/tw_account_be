package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	message_error "account-service/messageError"
	"account-service/model"
	"account-service/repository"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type registerRepository struct {
	db              *gorm.DB
	registerUtils   utils.RegisterUtils
	credentialUtils utils.CredentialUtils
}

func (r *registerRepository) CheckExistEmail(email string) (ok bool, err error) {
	var creadential *model.Credential

	errCredential := r.db.Model(&model.Credential{}).Where("email = ?", email).First(&creadential).Error

	if errCredential != nil && errCredential.Error() != message_error.RECORD_NOT_FOUND {
		return false, errCredential
	}

	if creadential.Id != 0 {
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

func (r *registerRepository) CreateProfile(saveInfo model.SaveRegister) (profile *model.Profile, err error) {
	var profileCreate = &model.Profile{
		Name: saveInfo.Username,
	}

	errCreate := r.db.Model(&model.Profile{}).Create(&profileCreate).Error

	if errCreate != nil {
		return nil, errCreate
	}

	return profileCreate, nil
}

func (r *registerRepository) CreateCredential(saveInfo *model.SaveRegister, profile *model.Profile, roleId uint) (err error) {
	var newCredential = &model.Credential{
		ProfileId: profile.Id,
		Username:  saveInfo.Username,
		Email:     saveInfo.Email,
		Password:  saveInfo.Password,
		RoleId:    roleId,
	}

	errCreate := r.db.Model(&model.Credential{}).Create(&newCredential).Error
	return errCreate
}

func (r *registerRepository) GetSaveInfo(saveInfoId uint) (saveInfo *model.SaveRegister, err error) {
	var saveInfoConfirm *model.SaveRegister

	errFind := r.db.Model(&model.SaveRegister{}).Where("id = ?", saveInfoId).First(&saveInfoConfirm).Error
	if errFind != nil {
		return nil, errFind
	}

	inValidTime := time.Now().After(saveInfoConfirm.FinishAt)
	if inValidTime {
		return nil, errors.New(message_error.CODE_EXPIRED)
	}

	return saveInfoConfirm, nil
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
