package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
