package response

import "time"

type SaveInfoRegisterResponse struct {
	Id       uint
	Email    string
	StartAt  time.Time `json:"startAt"`
	FinishAt time.Time `json:"finishAt"`
}
