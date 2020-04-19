package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCategoryById(categoryId int64) (*model.Category, error) {
	category, err := factories.FindCategoryById(categoryId)

	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category does not exists")
	}

	return category, nil
}
