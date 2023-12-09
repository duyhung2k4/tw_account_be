package utils

type TokenUtils interface {
	CreateToken(data map[string]interface{}) (token string, err error)
}
