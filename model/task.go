package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id        uint      `json:"id"`
	CreaterId uint      `json:"createrId"`
	ProjectId uint      `json:"projectId"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	StartAt   time.Time `json:"startAt"`
	FinishAt  time.Time `json:"finishAt"`
	Status    STATUS    `json:"status"`
	Detail    string    `json:"detail"`

	Creater *Profile `json:"creater" gorm:"foreign:CreaterId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Project *Project `json:"project" gorm:"foreign:ProjectId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime:true"`
}

type STATUS string

const (
	OPEN            STATUS = "open"
	CLOSE           STATUS = "close"
	IN_PROCESS      STATUS = "in_process"
	CANCELED        STATUS = "canceled"
	WILL_BE_CHECKED STATUS = "will_be_checked"
	LOOKING_BACK    STATUS = "looking_back"
)
