package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetProductItemById(productItemId int64) (*model.ProductItem, error) {
	productItem, err := factories.FindProductItemById(productItemId)

	if err != nil {
		return nil, err
	}

	if productItem == nil {
		return nil, errors.New("product_item does not exists")
	}

	return productItem, nil
}
