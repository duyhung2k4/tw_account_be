package request

import "account-service/model"

type RegisterRequest struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Email    string     `json:"email"`
	Role     model.ROLE `json:"role"`
}
