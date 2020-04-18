package model

type Order struct {
	Id                int64  `json:"_id"`
	Session           string `json:"session"`
	Address           string `json:"address"`
	PaymentStatus     string `json:"payment_status"`
	FulfillmentStatus string `json:"fulfillment_status"`
	Note              string `json:"note"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	PaidAt            string `json:"paid_at"`
	CancelledAt       string `json:"cancelled_at"`
}
