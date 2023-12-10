package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type ProjectRepository interface {
	GetProjectById(id uint) (projects *model.Project, err error)
	GetProjectByCreaterId(id uint) (projects []model.Project, err error)
	GetProjectJoined(id uint) (projects []model.Project, err error)
	CreateProject(req request.NewProjectRequest) (project *model.Project, err error)
	DeleteProject(id uint) (err error)
}
