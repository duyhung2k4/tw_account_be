package utils

type CredentialUtils interface {
	HashPassword(password string) (hashPassword string, err error)
	ComparePassword(password, hashPassword string) (success bool)
}
