package impl_controller

import (
	"account-service/controller"
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/service"
	impl_service "account-service/service/implService"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type accountController struct {
	accountService service.AccountService
}

func (a *accountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	var req request.FindAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	listAccount, errFind := a.accountService.FindAccount(req)
	if errFind != nil {
		response.ServerError(w, r, errFind)
		return
	}

	res := response.Response{
		Data:    listAccount,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func (a *accountController) AddUserToProject(w http.ResponseWriter, r *http.Request) {
	var req request.AddAccountToProject
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	newProjectProfile, errProjectProfile := a.accountService.AddAccountToProject(req)
	if errProjectProfile != nil {
		response.ServerError(w, r, errProjectProfile)
		return
	}

	res := response.Response{
		Data:    newProjectProfile,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func (a *accountController) GetUserProject(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	projectId64, errProjectId := strconv.ParseUint(idParam, 10, 64)

	if errProjectId != nil {
		response.BadRequest(w, r, errProjectId)
		return
	}

	projectId := uint(projectId64)
	listCredential, errListCredential := a.accountService.GetUserProject(projectId)
	if errListCredential != nil {
		response.ServerError(w, r, errListCredential)
		return
	}

	res := response.Response{
		Data:    listCredential,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func AccountControllerInit() controller.AccountController {
	return &accountController{
		accountService: impl_service.AccountServiceInit(),
	}
}
