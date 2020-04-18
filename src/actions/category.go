package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListCategory(queryCategory *model.QueryCategory) (interface{}, error) {
	query, err := queryCategory.GetQueryList()

	if err != nil {
		return nil, err
	}

	data, err := factories.FindListCategory()

	if err != nil {
		return nil, err
	}

	total, err := factories.CountCategory(query)

	if err != nil {
		return nil, err
	}

	return &response.DataList{
		Data:  data,
		Total: total,
		Page:  0,
		Limit: 0,
	}, nil
}

func GetOneCategory(categoryId *int64) (*model.Category, error) {

	return factories.FindCategoryById(categoryId)
}

func InsertCategory(insertCategory *model.BodyCategory) (*int64, error) {

	return factories.InsertCategory(insertCategory)
}

func GetAllCategoryTree() ([]*model.CategoryTree, error) {

	return factories.FindAllCategoryTree()
}
