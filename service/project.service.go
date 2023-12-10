package service

import (
	"account-service/dto/request"
	"account-service/model"
)

type ProjectService interface {
	GetProjectByCreaterId(id uint) (projects []model.Project, err error)
	GetProjectJoined(id uint) (projects []model.Project, err error)
	GetProjectById(id uint) (project *model.Project, err error)
	CreateProject(req request.NewProjectRequest) (project *model.Project, err error)
	DeleteProject(req request.DeleteProjectRequest) (err error)
}
