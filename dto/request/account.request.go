package request

type FindAccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AddAccountToProject struct {
	ProjectId        uint `json:"projectId"`
	CreaterProjectId uint `json:"createrProjectId"`
	JoinedId         uint `json:"joinedId"`
}
