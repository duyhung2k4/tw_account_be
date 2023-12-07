package controller

import "net/http"

type RegisterCotroller interface {
	SendInfoRegister(w http.ResponseWriter, r *http.Request)
	ConfirmCodeRegister(w http.ResponseWriter, r *http.Request)
}
