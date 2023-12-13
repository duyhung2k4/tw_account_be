package controller

import "net/http"

type TaskController interface {
	GetTaskOfProject(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateStatusTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	AddUserToTask(w http.ResponseWriter, r *http.Request)
	RemoveUserToTask(w http.ResponseWriter, r *http.Request)
	GetAllUserOfTask(w http.ResponseWriter, r *http.Request)
}
