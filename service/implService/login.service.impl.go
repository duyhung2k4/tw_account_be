package impl_service

import (
	"account-service/dto/request"
	"account-service/model"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
)

type loginService struct {
	loginRepo repository.LoginRepository
}

func (l *loginService) CheckCredential(info request.LoginRequest) (credential *model.Credential, err error) {
	newCredential, errCredential := l.loginRepo.FindCredential(info)

	return newCredential, errCredential
}

func LoginServiceInit() service.LoginService {
	return &loginService{
		loginRepo: impl_repository.LoginRepositoryInit(),
	}
}
