package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	"account-service/dto/response"
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

	sql := ""

	if req.Username != "" {
		sql = "username LIKE '" + req.Username + "%'"
	}

	if req.Email != "" {
		sql = "email LIKE '" + req.Email + "%'"
	}

	if sql == "" {
		return listCredential, nil
	}

	errListCredential := a.db.
		Model(&model.Credential{}).
		Preload("Profile").
		Preload("Role").
		Where(sql).
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

func (a *accountRepository) GetUserProject(projectId uint) (userOfProjects []response.UserOfProject, err error) {
	var listUserOfProjects []response.UserOfProject
	var listCredential []model.Credential
	var allCredential []model.Credential
	var listProjectProfile []model.ProjectProfile

	errProjectProfile := a.db.
		Model(&model.ProjectProfile{}).
		Preload("Profile").
		Where("project_id = ?", projectId).
		Find(&listProjectProfile).Error
	if errProjectProfile != nil {
		return listUserOfProjects, errProjectProfile
	}

	errAllCredential := a.db.
		Model(&model.Credential{}).
		Preload("Profile").
		Preload("Role").
		Find(&allCredential).Error
	if errAllCredential != nil {
		return listUserOfProjects, errAllCredential
	}

	for _, credential := range allCredential {
		var cre model.Credential
		var p model.ProjectProfile
		for _, projectProfile := range listProjectProfile {
			if projectProfile.ProfileId == credential.ProfileId {
				cre = credential
				p = projectProfile
				break
			}
		}

		if cre.Id == 0 || p.Id == 0 {
			continue
		}

		checkExist := 0
		for _, creInListReturn := range listCredential {
			if creInListReturn.Id == cre.Id {
				checkExist++
				break
			}
		}

		if checkExist > 0 {
			continue
		}

		listUserOfProjects = append(listUserOfProjects, response.UserOfProject{
			Credential:     cre,
			ProjectProfile: p,
		})
	}

	return listUserOfProjects, nil
}

func AccountRepositoryInit() repository.AccountRepository {
	return &accountRepository{
		db: config.GetDB(),
	}
}
