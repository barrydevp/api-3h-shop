package actions

import (
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetAllCategoryTree() ([]*model.CategoryTree, error) {

	return factories.FindAllCategoryTree()
}
