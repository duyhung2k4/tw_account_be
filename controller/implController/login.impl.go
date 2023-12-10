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
	"github.com/google/uuid"
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
		"id":       credential.Id,
		"email":    credential.Email,
		"userName": credential.Username,
		"role":     credential.Role.Code,
		"uuid":     uuid.New().String(),
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
		Data: response.LoginReponse{
			AccessToken:  token,
			RefreshToken: token,
			Credential:   *credential,
		},
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

func (l *loginController) LoginToken(w http.ResponseWriter, r *http.Request) {
	token := l.tokenUtils.GetToken(r)
	data, err := l.tokenUtils.ConvertToMap(token)
	fullDomain := r.Header.Get("Origin")

	if err != nil {
		response.ServerError(w, r, err)
		return
	}

	credentialId := uint(data["id"].(float64))
	credential, errCredential := l.loginService.FindCredential(credentialId)
	if errCredential != nil {
		response.ServerError(w, r, errCredential)
		return
	}

	newToken, errToken := l.tokenUtils.CreateToken(map[string]interface{}{
		"id":       credential.Id,
		"email":    credential.Email,
		"userName": credential.Username,
		"role":     credential.Role.Code,
		"uuid":     uuid.New().String(),
	})
	if errToken != nil {
		response.ServerError(w, r, errToken)
		return
	}

	accessCookie := http.Cookie{
		Name:    constant.ACCESS_TOKEN,
		Value:   newToken,
		Expires: time.Now().Add(3 * 24 * time.Hour),
		Domain:  fullDomain,
		Path:    "/",
	}

	refreshCookie := http.Cookie{
		Name:    constant.REFRESH_TOKEN,
		Value:   newToken,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		Domain:  fullDomain,
		Path:    "/",
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)

	res := response.Response{
		Data: response.LoginReponse{
			AccessToken:  newToken,
			RefreshToken: newToken,
			Credential:   *credential,
		},
		Message: "OK",
		Error:   "",
		Success: true,
	}

	render.JSON(w, r, res)
}

func LoginControllerInit() controller.LoginController {
	return &loginController{
		loginService: impl_service.LoginServiceInit(),
		tokenUtils:   impl_utils.TokenUtilsInit(),
	}
}
