package request

type DeleteProjectRequest struct {
	Id uint `json:"id"`
}

type NewProjectRequest struct {
	CreaterId uint   `json:"createrId"`
	Name      string `json:"name"`
}
