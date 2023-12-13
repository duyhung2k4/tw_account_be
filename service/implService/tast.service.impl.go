package impl_service

import (
	"account-service/dto/request"
	"account-service/model"
	"account-service/repository"
	impl_repository "account-service/repository/implRepository"
	"account-service/service"
)

type taskService struct {
	taskRepo repository.TaskRepository
}

func (t *taskService) GetTaskOfProject(projectId uint) (tasks []model.Task, err error) {
	listTask, errlistTask := t.taskRepo.GetTaskOfProject(projectId)
	return listTask, errlistTask
}

func (t *taskService) CreateTask(task request.CreateTaskRequest) (newTask *model.Task, err error) {
	taskCreate, errTaskCreate := t.taskRepo.CreateTask(task)
	return taskCreate, errTaskCreate
}

func (t *taskService) DeleteTask(taskId uint) (err error) {
	errDelete := t.taskRepo.DeleteTask(taskId)
	return errDelete
}

func (t *taskService) UpdateStatusTask(req model.Task) (newTask *model.Task, err error) {
	tastUpdate, errUpdate := t.taskRepo.UpdateStatusTask(req)
	return tastUpdate, errUpdate
}

func (t *taskService) AddUserToTask(req request.AddUserToTaskRequest) (taskProfile *model.TaskProfile, err error) {
	newTaskProfile, errCreate := t.taskRepo.AddUserToTask(req.CredentialId, req.TaskId)
	return newTaskProfile, errCreate
}

func (t *taskService) RemoveUserToTask(req request.RemoveUserToTaskRequest) (err error) {
	errDelete := t.taskRepo.RemoveUserToTask(req.TaskProfileId, req.CredentialId, req.TaskId)
	return errDelete
}

func TaskServiceInit() service.TaskService {
	return &taskService{
		taskRepo: impl_repository.TaskRepositoryInit(),
	}
}
