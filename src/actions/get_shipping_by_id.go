package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetShippingById(shippingId int64) (*model.Shipping, error) {
	shipping, err := factories.FindShippingById(shippingId)

	if err != nil {
		return nil, err
	}

	if shipping == nil {
		return nil, errors.New("shipping does not exists")
	}

	return shipping, nil
}
