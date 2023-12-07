package model

import (
	"time"

	"gorm.io/gorm"
)

type SaveRegister struct {
	Id       uint      `json:"id" gorm:"primaryKey"`
	Code     string    `json:"code"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     ROLE      `json:"role"`
	StartAt  time.Time `json:"startAt"`
	FinishAt time.Time `json:"finishAt"`

	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
