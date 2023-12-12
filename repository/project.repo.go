package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type ProjectRepository interface {
	GetProjectCreaterById(id uint, credentialId uint) (projects *model.Project, err error)
	GetProjectByCreaterId(id uint) (projects []model.Project, err error)
	GetProjectJoined(id uint) (projects []model.Project, err error)
	GetProjectJoinedById(id uint, credentialId uint) (project *model.Project, err error)
	CheckProjectOfCredential(credentialId uint, projectId uint) (ok bool, err error)
	CreateProject(req request.NewProjectRequest) (project *model.Project, err error)
	DeleteProject(req request.DeleteProjectRequest) (err error)
}
