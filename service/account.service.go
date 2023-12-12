package service

import (
	"account-service/dto/request"
	"account-service/model"
)

type AccountService interface {
	FindAccount(req request.FindAccountRequest) (credentials []model.Credential, err error)
	AddAccountToProject(req request.AddAccountToProject) (projectProfile *model.ProjectProfile, err error)
}
