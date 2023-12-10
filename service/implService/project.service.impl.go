package impl_service

import (
	"account-service/dto/request"
	"account-service/model"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
)

type projectService struct {
	projectRepo repository.ProjectRepository
}

func (p *projectService) GetProjectByCreaterId(id uint) (projects []model.Project, err error) {
	listProject, errProject := p.projectRepo.GetProjectByCreaterId(id)
	return listProject, errProject
}

func (p *projectService) GetProjectJoined(id uint) (projects []model.Project, err error) {
	listProject, errProject := p.projectRepo.GetProjectJoined(id)
	return listProject, errProject
}

func (p *projectService) GetProjectById(id uint) (project *model.Project, err error) {
	simpleProject, errProject := p.projectRepo.GetProjectById(id)
	return simpleProject, errProject
}

func (p *projectService) CreateProject(req request.NewProjectRequest) (project *model.Project, err error) {
	newProject, errCreate := p.projectRepo.CreateProject(req)
	return newProject, errCreate
}

func (p *projectService) DeleteProject(req request.DeleteProjectRequest) (err error) {
	errDelete := p.projectRepo.DeleteProject(req.Id)
	return errDelete
}

func ProjectServiceInit() service.ProjectService {
	return &projectService{
		projectRepo: impl_repository.ProjectRepositoryInit(),
	}
}
