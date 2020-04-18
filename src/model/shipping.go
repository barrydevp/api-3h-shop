package model

type Shipping struct {
	Id          int64  `json:"_id"`
	Carrier     string `json:"carrier"`
	Status      string `json:"status"`
	OrderId     int64  `json:"order_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeliveredAt string `json:"delivered_at"`
}
