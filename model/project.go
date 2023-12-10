package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	CreaterId uint   `json:"createrId"`
	Name      string `json:"name"`
	Code      string `json:"code" gorm:"unique"`

	Creater *Profile `json:"creater" gorm:"foreign:CreaterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
