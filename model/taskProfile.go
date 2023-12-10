package model

import (
	"time"

	"gorm.io/gorm"
)

type TaskProfile struct {
	Id            uint `json:"id"`
	TaskId        uint `json:"taskId"`
	ImplementerId uint `json:"implementerId"`

	Task        *Task    `json:"task" gorm:"foreign:TaskId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Implementer *Profile `json:"implementer" gorm:"foreign:ImplementerId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}
