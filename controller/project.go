package controller

import "net/http"

type ProjectController interface {
	GetProjectByCreaterId(w http.ResponseWriter, r *http.Request)
	// GetProjectJoined(w http.ResponseWriter, r *http.Request)
	GetProjectById(w http.ResponseWriter, r *http.Request)
}
