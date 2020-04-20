package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetOrderItemById(orderItemId int64) (*model.OrderItem, error) {
	orderItem, err := factories.FindOrderItemById(orderItemId)

	if err != nil {
		return nil, err
	}

	if orderItem == nil {
		return nil, errors.New("order_item does not exists")
	}

	return orderItem, nil
}
