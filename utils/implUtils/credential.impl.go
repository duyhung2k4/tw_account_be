package impl_utils

import (
	"account-service/config"
	"account-service/model"
	"account-service/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type credentialUtils struct {
	db *gorm.DB
}

func (c *credentialUtils) HashPassword(password string) (hashPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (c *credentialUtils) ComparePassword(password, passwordHash string) (success bool) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func (c *credentialUtils) GetRole() (roles []model.Role, err error) {
	var listRole []model.Role

	errFind := c.db.Model(&model.Role{}).Find(&listRole).Error

	return listRole, errFind
}

func CredentialUtilsInit() utils.CredentialUtils {
	return &credentialUtils{
		db: config.GetDB(),
	}
}
