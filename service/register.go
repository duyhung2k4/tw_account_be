package service

import (
	"account-service/dto/request"
	"account-service/dto/response"
)

type RegisterService interface {
	HandleSendInforegister(info request.RegisterRequest) (res *response.SaveInfoRegisterResponse, err error)
}
