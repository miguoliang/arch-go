package dto

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type GroupList ListResponse[Group]
