package service

import (
	"account-service/dto/request"
	"account-service/model"
)

type TaskService interface {
	GetTaskOfProject(projectId uint) (tasks []model.Task, err error)
	CreateTask(task request.CreateTaskRequest) (newTask *model.Task, err error)
	DeleteTask(taskId uint) (err error)
	UpdateStatusTask(req model.Task) (newTask *model.Task, err error)
	AddUserToTask(req request.AddUserToTaskRequest) (taskProfile *model.TaskProfile, err error)
	RemoveUserToTask(req request.RemoveUserToTaskRequest) (err error)
}
