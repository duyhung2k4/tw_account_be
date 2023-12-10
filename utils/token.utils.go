package utils

import "net/http"

type TokenUtils interface {
	CreateToken(data map[string]interface{}) (token string, err error)
	GetToken(r *http.Request) (token string)
	ConvertToMap(token string) (data map[string]interface{}, err error)
}
