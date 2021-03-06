package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetProductById(productItemId int64) (*model.Product, error) {
	productItem, err := factories.FindProductById(productItemId)

	if err != nil {
		return nil, err
	}

	if productItem == nil {
		return nil, errors.New("product does not exists")
	}

	return productItem, nil
}
