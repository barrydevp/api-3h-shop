package model

type Product struct {
	Id          int64  `json:"_id"`
	CategoryId  int64  `json:"category_id"`
	Name        string `json:"name"`
	OutPrice    string `json:"out_price"`
	Discount    string `json:"discount"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
