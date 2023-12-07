package impl_service

import (
	"account-service/dto/request"
	"account-service/dto/response"
	message_error "account-service/messageError"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"errors"
)

type registerService struct {
	registerRepository repository.RegisterRepository
	emailUtils         utils.EmailUtils
}

func (r *registerService) HandleSendInforegister(info request.RegisterRequest) (res *response.SaveInfoRegisterResponse, err error) {

	ok, errCheckEmail := r.registerRepository.CheckExistEmail(info.Email)

	if errCheckEmail != nil {
		return nil, errCheckEmail
	}

	if !ok {
		return nil, errors.New(message_error.EMAIL_EXIST)
	}

	newInfo, errCreateInfo := r.registerRepository.CreateSaveInfoRegister(info)

	if errCreateInfo != nil {
		return nil, errCreateInfo
	}

	errSendCodeEmail := r.emailUtils.SendEmailConfirmCodeRegister(newInfo.Email, newInfo.Code)

	if errSendCodeEmail != nil {
		return nil, errSendCodeEmail
	}

	saveInfoRegister := response.SaveInfoRegisterResponse{
		Id:       newInfo.Id,
		Email:    newInfo.Email,
		StartAt:  newInfo.StartAt,
		FinishAt: newInfo.FinishAt,
	}

	return &saveInfoRegister, nil
}

func RegisterServiceInit() service.RegisterService {
	return &registerService{
		registerRepository: impl_repository.RegisterRepositoryInit(),
		emailUtils:         impl_utils.EmailUtilsInit(),
	}
}
