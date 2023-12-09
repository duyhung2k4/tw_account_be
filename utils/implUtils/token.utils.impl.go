package impl_utils

import (
	"account-service/config"
	"account-service/utils"

	"github.com/go-chi/jwtauth/v5"
)

type tokenUtils struct {
	jwt *jwtauth.JWTAuth
}

func (t *tokenUtils) CreateToken(data map[string]interface{}) (token string, err error) {
	_, tokenString, errToken := t.jwt.Encode(data)
	return tokenString, errToken
}

func TokenUtilsInit() utils.TokenUtils {
	return &tokenUtils{
		jwt: config.GetJWT(),
	}
}
