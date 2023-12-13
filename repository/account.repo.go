package repository

import (
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/model"
)

type AccountRepository interface {
	FindAccount(req request.FindAccountRequest) (creadential []model.Credential, err error)
	GetUserProject(projectId uint) (userOfProjects []response.UserOfProject, err error)
	AddAccountToProject(accountId uint, projectId uint) (projectProfile *model.ProjectProfile, err error)
}
