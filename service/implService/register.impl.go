package impl_service

import (
	"account-service/config"
	"account-service/dto/request"
	"account-service/dto/response"
	message_error "account-service/messageError"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"errors"

	"gorm.io/gorm"
)

type registerService struct {
	db                 *gorm.DB
	registerRepository repository.RegisterRepository
	emailUtils         utils.EmailUtils
	credentialUtils    utils.CredentialUtils
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
		Username: newInfo.Username,
		Password: info.Password,
		Role:     info.Role,
		StartAt:  newInfo.StartAt,
		FinishAt: newInfo.FinishAt,
	}

	return &saveInfoRegister, nil
}

func (r *registerService) HandleConfirmCode(confirmInfo request.ConfirmInfo) (err error) {

	saveInfo, errSaveInfo := r.registerRepository.GetSaveInfo(confirmInfo.SaveInfoId)
	if errSaveInfo != nil {
		return errSaveInfo
	}

	profile, errCreateProfile := r.registerRepository.CreateProfile(*saveInfo)
	if errCreateProfile != nil {
		return errCreateProfile
	}

	roles, errFindRole := r.credentialUtils.GetRole()
	if errFindRole != nil {
		return errFindRole
	}

	var roleIdCredentile *uint
	for _, role := range roles {
		if role.Code == saveInfo.Role {
			roleIdCredentile = &role.Id
			break
		}
	}
	if roleIdCredentile == nil {
		return errors.New(message_error.INVALID_ROLE)
	}

	errCreateCredential := r.registerRepository.CreateCredential(saveInfo, profile, *roleIdCredentile)

	return errCreateCredential
}

func RegisterServiceInit() service.RegisterService {
	return &registerService{
		db:                 config.GetDB(),
		registerRepository: impl_repository.RegisterRepositoryInit(),
		emailUtils:         impl_utils.EmailUtilsInit(),
		credentialUtils:    impl_utils.CredentialUtilsInit(),
	}
}
