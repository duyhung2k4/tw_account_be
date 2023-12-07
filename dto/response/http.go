package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

type MetaResonse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Page    Page        `json:"page"`
}

type Page struct {
	Page     int  `json:"page"`
	PageSize int  `json:"pageSize"`
	Total    uint `json:"total"`
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	res := Response{
		Data:    nil,
		Success: false,
		Message: "Bad request!",
		Error:   err.Error(),
	}
	render.JSON(w, r, res)
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	res := MetaResonse{
		Data:    nil,
		Success: false,
		Message: err.Error(),
		Error:   err.Error(),
		Page: Page{
			Page:     0,
			PageSize: 0,
			Total:    0,
		},
	}
	render.JSON(w, r, res)
}
