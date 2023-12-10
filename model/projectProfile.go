package model

import (
	"time"

	"gorm.io/gorm"
)

type ProjectProfile struct {
	Id        uint `json:"id"`
	ProjectId uint `json:"projectId"`
	ProfileId uint `json:"profileId"`

	Project *Project `json:"project" gorm:"foreignKey:ProjectId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
