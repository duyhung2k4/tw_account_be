package controller

import "net/http"

type LoginController interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginToken(w http.ResponseWriter, r *http.Request)
}
