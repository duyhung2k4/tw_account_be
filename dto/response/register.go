package response

import (
	"account-service/model"
	"time"
)

type SaveInfoRegisterResponse struct {
	Id       uint       `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Username string     `json:"username"`
	Role     model.ROLE `json:"role"`
	StartAt  time.Time  `json:"startAt"`
	FinishAt time.Time  `json:"finishAt"`
}
