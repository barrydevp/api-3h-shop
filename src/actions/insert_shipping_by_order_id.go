package actions

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertShippingByOrderId(orderId int64, body *model.BodyShipping) (*model.Shipping, error) {
	body.OrderId = &orderId

	return InsertShipping(body)
}
