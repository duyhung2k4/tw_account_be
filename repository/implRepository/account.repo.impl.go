package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	message_error "account-service/messageError"
	"account-service/model"
	"account-service/repository"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func (a *accountRepository) FindAccount(req request.FindAccountRequest) (creadential []model.Credential, err error) {
	var listCredential []model.Credential

	errListCredential := a.db.
		Model(&model.Credential{}).
		Where("username LIKE ?% OR email LIKE ?%", req.Username, req.Email).
		Find(&listCredential).Error
	if errListCredential != nil && errListCredential.Error() != message_error.RECORD_NOT_FOUND {
		return listCredential, errListCredential
	}

	return listCredential, nil
}

func (a *accountRepository) AddAccountToProject(accountId uint, projectId uint) (projectProfile *model.ProjectProfile, err error) {
	var newProjectProfile = &model.ProjectProfile{
		ProjectId: projectId,
		ProfileId: accountId,
	}

	errCreate := a.db.Model(&model.ProjectProfile{}).Create(&newProjectProfile).Error
	return newProjectProfile, errCreate
}

func AccountRepositoryInit() repository.AccountRepository {
	return &accountRepository{
		db: config.GetDB(),
	}
}
