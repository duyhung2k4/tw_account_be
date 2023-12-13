package repository

import (
	"account-service/dto/request"
	"account-service/model"
)

type TaskRepository interface {
	GetTaskOfProject(projectId uint) (task []model.Task, err error)
	CreateTask(task request.CreateTaskRequest) (newTask *model.Task, err error)
	DeleteTask(taskId uint) (err error)
	UpdateStatusTask(req model.Task) (newTask *model.Task, err error)

	GetAllUserOfTask() (taskProfiles []model.TaskProfile, err error)
	AddUserToTask(credentialId uint, taskId uint) (taskProfile *model.TaskProfile, err error)
	RemoveUserToTask(credentialId uint, taskId uint) (err error)
}
