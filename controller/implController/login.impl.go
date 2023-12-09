package impl_controller

import (
	"account-service/constant"
	"account-service/controller"
	"account-service/dto/request"
	"account-service/dto/response"
	"account-service/service"
	impl_service "account-service/service/implService"
	"account-service/utils"
	impl_utils "account-service/utils/implUtils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type loginController struct {
	loginService service.LoginService
	tokenUtils   utils.TokenUtils
}

// @Summary      Send info login
// @Description  Send info login
// @Tags         Login
// @Accept       json
// @Produce 		 json
// @Param        req body request.LoginRequest true "Send info"
// @Success      200  {object}  response.Response
// @Router       /public/login [post]
func (l *loginController) Login(w http.ResponseWriter, r *http.Request) {
	var info request.LoginRequest
	errRequest := json.NewDecoder(r.Body).Decode(&info)
	fullDomain := r.Header.Get("Origin")

	if errRequest != nil {
		response.BadRequest(w, r, errRequest)
		return
	}

	credential, errCredential := l.loginService.CheckCredential(info)
	if errCredential != nil {
		response.ServerError(w, r, errCredential)
		return
	}

	token, errToken := l.tokenUtils.CreateToken(map[string]interface{}{
		"id":        credential.Id,
		"email":     credential.Email,
		"user_name": credential.Username,
		"role":      credential.Role.Code,
	})
	if errToken != nil {
		response.ServerError(w, r, errToken)
		return
	}

	accessCookie := http.Cookie{
		Name:    constant.ACCESS_TOKEN,
		Value:   token,
		Expires: time.Now().Add(3 * 24 * time.Hour),
		Domain:  fullDomain,
		Path:    "/",
	}

	refreshCookie := http.Cookie{
		Name:    constant.REFRESH_TOKEN,
		Value:   token,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		Domain:  fullDomain,
		Path:    "/",
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)

	res := response.Response{
		Data:    "",
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

func (l *loginController) LoginToken(w http.ResponseWriter, r *http.Request) {

}

func LoginControllerInit() controller.LoginController {
	return &loginController{
		loginService: impl_service.LoginServiceInit(),
		tokenUtils:   impl_utils.TokenUtilsInit(),
	}
}
