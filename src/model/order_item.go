package model

type OrderItem struct {
	Id            int64  `json:"_id"`
	ProductId     int64  `json:"product_id"`
	ProductItemId int64  `json:"product_item_id"`
	OrderId       int64  `json:"order_id"`
	Quantity      int64  `json:"quantity"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
