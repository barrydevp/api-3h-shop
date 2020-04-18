package model

type User struct {
	Id        int64  `json:"_id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password,omitempty"`
	Address   string `json:"address,omitempty"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
