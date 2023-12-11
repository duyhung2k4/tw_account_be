package impl_controller

import (
	"account-service/controller"
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/service"
	impl_service "account-service/service/implService"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// @Summary      Get project create by createrId
// @Description  Get project create by createrId
// @Tags         Project
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Project id"
// @Success      200  {object}  response.Response
// @Router       /protected/project/creater_id_detail/{id} [get]
func (p *projectController) GetProjectCreaterById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	projectId64, errProjectId := strconv.ParseUint(idParam, 10, 64)
	if errProjectId != nil {
		response.BadRequest(w, r, errProjectId)
		return
	}
	projectId := uint(projectId64)

	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)
	if errToken != nil {
		response.BadRequest(w, r, errToken)
		return
	}

	credentialId := uint(mapData["id"].(float64))
	project, errProject := p.projectService.GetProjectCreaterById(projectId, credentialId)
	if errProject != nil {
		response.ServerError(w, r, errProject)
		return
	}

	res := response.Response{
		Data:    project,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

// @Summary      Get project joined
// @Description  Get project joined
// @Tags         Project
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Router       /protected/project/joined [get]
func (p *projectController) GetProjectJoined(w http.ResponseWriter, r *http.Request) {
	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)

	if errToken != nil {
		response.ServerError(w, r, errToken)
		return
	}

	credentialId := uint(mapData["id"].(float64))
	projects, errProjects := p.projectService.GetProjectJoined(credentialId)
	if errProjects != nil {
		response.ServerError(w, r, errProjects)
		return
	}

	res := response.Response{
		Data:    projects,
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

// @Summary      Get project joined detail
// @Description  Get project joined detail
// @Tags         Project
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Project id"
// @Success      200  {object}  response.Response
// @Router       /protected/project/joined_detail/{id} [get]
func (p *projectController) GetProjectJoinedById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	projectId64, errProjectId := strconv.ParseUint(idParam, 10, 64)
	if errProjectId != nil {
		response.BadRequest(w, r, errProjectId)
		return
	}
	projectId := uint(projectId64)

	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)
	if errToken != nil {
		response.BadRequest(w, r, errToken)
		return
	}

	credentialId := uint(mapData["id"].(float64))
	project, errProject := p.projectService.GetProjectJoinedById(projectId, credentialId)
	if errProject != nil {
		response.ServerError(w, r, errProject)
		return
	}

	res := response.Response{
		Data:    project,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

// @Summary      Create Project
// @Description  Create Project
// @Tags         Project
// @Accept       json
// @Produce 		 json
// @Param        req body request.NewProjectRequest true "project"
// @Success      200  {object}  response.Response
// @Router       /protected/project/create [post]
func (p *projectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	var newProject request.NewProjectRequest
	errProject := json.NewDecoder(r.Body).Decode(&newProject)

	if errProject != nil {
		response.ServerError(w, r, errProject)
		return
	}

	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)
	if errToken != nil {
		response.ServerError(w, r, errProject)
		return
	}
	credentialId := uint(mapData["id"].(float64))

	newProject.CreaterId = credentialId

	newProjectReturn, errNewProjectReturn := p.projectService.CreateProject(newProject)
	if errNewProjectReturn != nil {
		response.ServerError(w, r, errNewProjectReturn)
		return
	}

	res := response.Response{
		Data:    newProjectReturn,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

// @Summary      Delete Project
// @Description  Delete Project
// @Tags         Project
// @Accept       json
// @Produce 		 json
// @Param        req body request.DeleteProjectRequest true "project"
// @Success      200  {object}  response.Response
// @Router       /protected/project/delete [delete]
func (p *projectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var deleteProject request.DeleteProjectRequest
	errProject := json.NewDecoder(r.Body).Decode(&deleteProject)

	if errProject != nil {
		response.ServerError(w, r, errProject)
		return
	}

	token := p.tokenUtils.GetToken(r)
	mapData, errToken := p.tokenUtils.ConvertToMap(token)
	if errToken != nil {
		response.ServerError(w, r, errProject)
		return
	}
	credentialId := uint(mapData["id"].(float64))
	deleteProject.CreaterId = credentialId

	errDelete := p.projectService.DeleteProject(deleteProject)
	if errDelete != nil {
		response.ServerError(w, r, errDelete)
		return
	}

	res := response.Response{
		Data:    nil,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func ProjectControllerInit() controller.ProjectController {
	return &projectController{
		tokenUtils:     impl_utils.TokenUtilsInit(),
		projectService: impl_service.ProjectServiceInit(),
	}
}
