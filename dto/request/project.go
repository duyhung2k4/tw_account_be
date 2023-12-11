package request

type DeleteProjectRequest struct {
	Id        uint `json:"id"`
	CreaterId uint `json:"createrId"`
}

type NewProjectRequest struct {
	CreaterId uint   `json:"createrId"`
	Name      string `json:"name"`
}
