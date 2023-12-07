package utils

import "time"

type RegisterUtils interface {
	CreateCode() (code string)
	CreateTimeExist() (startAt time.Time, finishAt time.Time)
}
