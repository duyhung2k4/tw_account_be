package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type TaskRepository interface {
	GetTaskOfProject(projectId uint) (task []model.Task, err error)
	CreateTask(task request.CreateTaskRequest) (newTask *model.Task, err error)
	DeleteTask(taskId uint) (err error)
	UpdateStatusTask(taskId uint, status model.STATUS) (newTask *model.Task, err error)
	AddUserToTask(credentialId uint, taskId uint) (taskProfile *model.TaskProfile, err error)
	RemoveUserToTask(taskProfileId uint, credentialId uint, taskId uint) (err error)
}
