package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetWarrantyById(warrantyId int64) (*model.Warranty, error) {
	warranty, err := factories.FindWarrantyById(warrantyId)

	if err != nil {
		return nil, err
	}

	if warranty == nil {
		return nil, errors.New("warranty does not exists")
	}

	return warranty, nil
}
