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

	"gorm.io/gorm"
)

type loginRepository struct {
	db             *gorm.DB
	credentialUtil utils.CredentialUtils
}

func (l *loginRepository) FindCredential(info request.LoginRequest) (credential *model.Credential, err error) {
	var newCredential *model.Credential

	errCredential := l.db.
		Model(&model.Credential{}).
		Preload("Role").
		Preload("Profile").
		Where("username = ?", info.Username).
		First(&newCredential).Error

	if err != nil {
		return nil, errCredential
	}

	if ok := l.credentialUtil.ComparePassword(info.Password, newCredential.Password); !ok {
		return nil, errors.New(message_error.ERROR_PASSWORD)
	}

	return newCredential, nil
}

func LoginRepositoryInit() repository.LoginRepository {
	return &loginRepository{
		db:             config.GetDB(),
		credentialUtil: impl_utils.CredentialUtilsInit(),
	}
}
