package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCustomerById(customerId int64) (*model.Customer, error) {
	customer, err := factories.FindCustomerById(customerId)

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("customer does not exists")
	}

	return customer, nil
}
