package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type AccountRepository interface {
	FindAccount(req request.FindAccountRequest) (creadential []model.Credential, err error)
	AddAccountToProject(accountId uint, projectId uint) (projectProfile *model.ProjectProfile, err error)
}
