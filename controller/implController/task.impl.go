package impl_controller

import (
	"account-service/controller"
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/model"
	"account-service/service"
	impl_service "account-service/service/implService"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type taskController struct {
	taskService service.TaskService
}

func (t *taskController) GetTaskOfProject(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	projectId64, errProjectId := strconv.ParseUint(idParam, 10, 64)
	if errProjectId != nil {
		response.BadRequest(w, r, errProjectId)
		return
	}
	projectId := uint(projectId64)

	listTask, errListTask := t.taskService.GetTaskOfProject(projectId)
	if errListTask != nil {
		response.ServerError(w, r, errListTask)
		return
	}

	res := response.Response{
		Data:    listTask,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func (t *taskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskCreate request.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&taskCreate)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	newTask, errCreate := t.taskService.CreateTask(taskCreate)
	if errCreate != nil {
		response.ServerError(w, r, errCreate)
		return
	}

	res := response.Response{
		Data:    newTask,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func (t *taskController) UpdateStatusTask(w http.ResponseWriter, r *http.Request) {
	var taskUpdate model.Task
	err := json.NewDecoder(r.Body).Decode(&taskUpdate)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	newTask, errCreate := t.taskService.UpdateStatusTask(taskUpdate)
	if errCreate != nil {
		response.ServerError(w, r, errCreate)
		return
	}

	res := response.Response{
		Data:    newTask,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func (t *taskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var deleteTask request.DeleteTaskRequest
	err := json.NewDecoder(r.Body).Decode(&deleteTask)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	errDelete := t.taskService.DeleteTask(deleteTask.TaskId)
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

func (t *taskController) AddUserToTask(w http.ResponseWriter, r *http.Request) {
	var addUserReq request.AddUserToTaskRequest
	err := json.NewDecoder(r.Body).Decode(&addUserReq)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	newUser, errUser := t.taskService.AddUserToTask(addUserReq)
	if errUser != nil {
		response.ServerError(w, r, errUser)
		return
	}

	res := response.Response{
		Data:    newUser,
		Message: "OK",
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)

}

func (t *taskController) RemoveUserToTask(w http.ResponseWriter, r *http.Request) {
	var removeUserReq request.RemoveUserToTaskRequest
	err := json.NewDecoder(r.Body).Decode(&removeUserReq)

	if err != nil {
		response.BadRequest(w, r, err)
		return
	}

	errRemove := t.taskService.RemoveUserToTask(removeUserReq)
	if errRemove != nil {
		response.ServerError(w, r, errRemove)
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

func TaskControllerInit() controller.TaskController {
	return &taskController{
		taskService: impl_service.TaskServiceInit(),
	}
}
