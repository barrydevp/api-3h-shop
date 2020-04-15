package model

type Category struct {
	Id        int    `json:"_id"`
	Name      string `json:"name"`
	ParentId  string `json:"parent_id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}
