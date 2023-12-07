package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Commune      string `json:"commune"`
	DistrictCode string `json:"districtCode"`
	ProvinceCode string `json:"provinceCode"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
