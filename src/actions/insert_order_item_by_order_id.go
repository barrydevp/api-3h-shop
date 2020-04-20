package actions

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertOrderItemByOrderId(orderId int64, body *model.BodyOrderItem) (*model.OrderItem, error) {
	body.OrderId = &orderId

	return InsertOrderItem(body)
}
