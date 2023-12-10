package impl_controller

import (
	"account-service/controller"
	"account-service/dto/response"
	"account-service/service"
	impl_service "account-service/service/implService"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type projectController struct {
	tokenUtils     utils.TokenUtils
	projectService service.ProjectService
}

// @Summary      Get project by createrId
// @Description  Get project by createrId
// @Tags         Project
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  response.Response
// @Router       /protected/project/creater_id [get]
func (p *projectController) GetProjectByCreaterId(w http.ResponseWriter, r *http.Request) {
	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)

	if errToken != nil {
		response.BadRequest(w, r, errToken)
		return
	}

	credentialId := uint(mapData["id"].(float64))
	projects, errProjects := p.projectService.GetProjectByCreaterId(credentialId)
	if errProjects != nil {
		response.ServerError(w, r, errProjects)
		return
	}

	res := response.Response{
		Data:    projects,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

// func (p *projectController) GetProjectJoined(w http.ResponseWriter, r *http.Request) {
// 	return
// }

func (p *projectController) GetProjectById(w http.ResponseWriter, r *http.Request) {
	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)

	if errToken != nil {
		response.BadRequest(w, r, errToken)
		return
	}

	credentialId := uint(mapData["id"].(float64))
	log.Println(credentialId)
	return
}

func ProjectControllerInit() controller.ProjectController {
	return &projectController{
		tokenUtils:     impl_utils.TokenUtilsInit(),
		projectService: impl_service.ProjectServiceInit(),
	}
}
