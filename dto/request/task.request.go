package request

import (
	"account-service/model"
	"time"
)

type CreateTaskRequest struct {
	CreaterId uint         `json:"createrId"`
	ProjectId uint         `json:"projectId"`
	Name      string       `json:"name"`
	Level     model.LEVEL  `json:"level"`
	StartAt   time.Time    `json:"startAt"`
	FinishAt  time.Time    `json:"finishAt"`
	Status    model.STATUS `json:"status"`
	Detail    string       `json:"detail"`
}

type DeleteTaskRequest struct {
	TaskId uint `json:"taskId"`
}

type UpdateStatusTaskRequest struct {
	TaskId uint         `json:"taskId"`
	Status model.STATUS `json:"status"`
}

type AddUserToTaskRequest struct {
	CredentialId uint `json:"credentialId"`
	TaskId       uint `json:"taskId"`
}

type RemoveUserToTaskRequest struct {
	CredentialId uint `json:"credentialId"`
	TaskId       uint `json:"taskId"`
}
