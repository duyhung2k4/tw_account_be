package impl_utils

import (
	"account-service/config"
	"account-service/utils"
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

type tokenUtils struct {
	jwt *jwtauth.JWTAuth
}

func (t *tokenUtils) CreateToken(data map[string]interface{}) (token string, err error) {
	_, tokenString, errToken := t.jwt.Encode(data)
	return tokenString, errToken
}

func (t *tokenUtils) GetToken(r *http.Request) (token string) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	return reqToken
}

func (t *tokenUtils) ConvertToMap(token string) (data map[string]interface{}, err error) {
	jwtToken, errToken := config.GetJWT().Decode(token)

	if errToken != nil {
		return map[string]interface{}{}, errToken
	}

	mapData, errMap := jwtToken.AsMap(context.Background())

	return mapData, errMap
}

func TokenUtilsInit() utils.TokenUtils {
	return &tokenUtils{
		jwt: config.GetJWT(),
	}
}
