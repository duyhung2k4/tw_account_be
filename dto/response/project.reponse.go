package response

import "account-service/model"

type UserOfProject struct {
	Credential     model.Credential     `json:"credential"`
	ProjectProfile model.ProjectProfile `json:"projectProfile"`
}
