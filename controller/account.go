package controller

import "net/http"

type AccountController interface {
	GetAccount(w http.ResponseWriter, r *http.Request)
	AddUserToProject(w http.ResponseWriter, r *http.Request)
}
