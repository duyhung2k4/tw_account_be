package impl_controller

import (
	"account-service/controller"
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/service"
	impl_service "account-service/service/implService"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type register struct {
	registerService service.RegisterService
}

// @Summary      Send info register
// @Description  Send info register
// @Tags         Register
// @Accept       json
// @Produce 		 json
// @Param        req body request.RegisterRequest true "Send info"
// @Success      200  {object}  response.Response
// @Router       /public/send_info [post]
func (re *register) SendInfoRegister(w http.ResponseWriter, r *http.Request) {
	var infoRegister request.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&infoRegister); err != nil {
		response.BadRequest(w, r, err)
		return
	}

	saveInfo, errSaveInfo := re.registerService.HandleSendInforegister(infoRegister)
	if errSaveInfo != nil {
		response.ServerError(w, r, errSaveInfo)
		return
	}

	res := response.Response{
		Data:    saveInfo,
		Message: "OK",
		Success: true,
	}

	render.JSON(w, r, res)
}

func (re *register) ConfirmCodeRegister(w http.ResponseWriter, r *http.Request) {}

func RegisterCotrollerInit() controller.RegisterCotroller {
	return &register{
		registerService: impl_service.RegisterServiceInit(),
	}
}
