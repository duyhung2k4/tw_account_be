package service

import (
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/model"
)

type AccountService interface {
	FindAccount(req request.FindAccountRequest) (credentials []model.Credential, err error)
	GetUserProject(projectId uint) (userOfProjects []response.UserOfProject, err error)
	AddAccountToProject(req request.AddAccountToProject) (projectProfile *model.ProjectProfile, err error)
}
