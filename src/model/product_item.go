package model

type ProductItem struct {
	Id        int64  `json:"_id"`
	ProductId int64  `json:"product_id"`
	Stock     string `json:"stock"`
	InPrice   string `json:"in_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	ExpiredAt string `json:"expired_at"`
}
