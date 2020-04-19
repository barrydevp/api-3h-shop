package actions

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/common/response"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListCategory(queryCategory *model.QueryCategory) (interface{}, error) {
	queryList, err := queryCategory.GetQueryList()
	queryCountList, err := queryCategory.GetQueryCountList()

	if err != nil {
		return nil, err
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindCategory(queryList)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		total, err := factories.CountCategory(queryCountList)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- total
	}()

	var data interface{}
	var total int

	for i := 0; i < 2; i++ {
		select {
		case val := <-resolveChan:
			tmp, ok := val.(int)
			if ok {
				total = tmp
			} else {
				data = val
			}

		case err := <-rejectChan:
			return nil, err
		}
	}

	return &response.DataList{
		Data:  data,
		Total: total,
		Page:  *queryCategory.Page,
		Limit: *queryCategory.Limit,
	}, nil
}

func GetCategoryById(categoryId int64) (*model.Category, error) {

	return factories.FindCategoryById(categoryId)
}

func InsertCategory(insertCategory *model.BodyCategory) (*int64, error) {

	return factories.InsertCategory(insertCategory)
}

func GetAllCategoryTree() ([]*model.CategoryTree, error) {

	return factories.FindAllCategoryTree()
}

func GetCategoryTreeById(categoryId int64) (*model.CategoryTree, error) {
	where := " parent_id=?"
	query := connect.QueryMySQL{
		Where: &where,
		Args:  []interface{}{categoryId},
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindCategory(&query)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		category, err := factories.FindCategoryById(categoryId)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- category
	}()

	var list []*model.Category
	var category *model.Category

	for i := 0; i < 2; i++ {
		select {
		case data := <-resolveChan:
			switch val := data.(type) {
			case *model.Category:
				category = val
			case []*model.Category:
				list = val
			}
		case err := <-rejectChan:
			return nil, err
		}
	}

	categoryTree := model.CategoryTree{
		Category: category,
	}

	if len(list) > 0 {
		_, categoryTree.Children = factories.GetCategoryChildren(list, *categoryTree.Id)
	}

	return &categoryTree, nil
}
