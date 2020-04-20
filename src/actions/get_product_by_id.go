package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetProductById(orderId int64) (*model.Product, error) {
	order, err := factories.FindProductById(orderId)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order does not exists")
	}

	return order, nil
}
