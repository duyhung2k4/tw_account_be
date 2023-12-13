package impl_service

import (
	"account-service/dto/request"
	"account-service/dto/response"
	message_error "account-service/messageError"
	"account-service/model"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
	"errors"
)

type accountService struct {
	accountRepo repository.AccountRepository
	projectRepo repository.ProjectRepository
}

func (a *accountService) FindAccount(req request.FindAccountRequest) (credentials []model.Credential, err error) {
	listCredential, errFind := a.accountRepo.FindAccount(req)
	return listCredential, errFind
}

func (a *accountService) AddAccountToProject(req request.AddAccountToProject) (projectProfile *model.ProjectProfile, err error) {
	ok, errCheckProject := a.projectRepo.CheckProjectOfCredential(req.CreaterProjectId, req.ProjectId)

	if errCheckProject != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New(message_error.NOT_IS_PROJECT)
	}

	newProjectProfle, errProjectProfile := a.accountRepo.AddAccountToProject(req.JoinedId, req.ProjectId)
	return newProjectProfle, errProjectProfile
}

func (a *accountService) GetUserProject(projectId uint) (userOfProjects []response.UserOfProject, err error) {
	listUserOfProject, errListUserOfProject := a.accountRepo.GetUserProject(projectId)
	return listUserOfProject, errListUserOfProject
}

func AccountServiceInit() service.AccountService {
	return &accountService{
		accountRepo: impl_repository.AccountRepositoryInit(),
		projectRepo: impl_repository.ProjectRepositoryInit(),
	}
}
