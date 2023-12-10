package response

import "account-service/model"

type LoginReponse struct {
	AccessToken  string           `json:"accessToken"`
	RefreshToken string           `json:"refreshToken"`
	Credential   model.Credential `json:"credential"`
}
