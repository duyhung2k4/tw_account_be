package controller

import "net/http"

type ProjectController interface {
	GetProjectByCreaterId(w http.ResponseWriter, r *http.Request)
	GetProjectCreaterById(w http.ResponseWriter, r *http.Request)
	GetProjectJoined(w http.ResponseWriter, r *http.Request)
	GetProjectJoinedById(w http.ResponseWriter, r *http.Request)
	CreateProject(w http.ResponseWriter, r *http.Request)
	DeleteProject(w http.ResponseWriter, r *http.Request)
}
