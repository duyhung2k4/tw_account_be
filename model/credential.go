package model

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	RoleId    uint   `json:"roleId"`
	ProfileId uint   `json:"profileId"`
	ShopId    uint   `json:"shopId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`

	Profile *Profile `json:"profile" gorm:"foreignKey:ProfileId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Role    *Role    `json:"role" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
