package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCategoryTreeById(categoryId int64) (*model.CategoryTree, error) {
	query := connect.QueryMySQL{
		QueryString: "WHERE parent_id=?",
		Args:        []interface{}{categoryId},
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
				if category == nil {
					return nil, errors.New("category does not exists")
				}
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
