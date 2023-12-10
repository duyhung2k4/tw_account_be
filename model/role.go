package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Code ROLE   `json:"code" gorm:"unique"`
	Name string `json:"name"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

type ROLE string

const (
	USER  ROLE = "user"
	ADMIN ROLE = "admin"
)
