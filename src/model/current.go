package model

type Current struct {
	Session *string
}

type CurrentResponse struct {
	Order     *Order     `json:"order"`
	Session   *string    `json:"session"`
}
