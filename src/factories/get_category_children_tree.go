package factories

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCategoryChildren(all []*model.Category, parentId int64) ([]*model.Category, []*model.CategoryTree) {
	var allCategoryTree []*model.CategoryTree
	var remainAllCategory []*model.Category

	for _, category := range all {
		if *category.ParentId == parentId {
			allCategoryTree = append(allCategoryTree, &model.CategoryTree{
				Category: category,
				Children: nil,
			})
		} else {
			remainAllCategory = append(remainAllCategory, category)
		}
	}

	for _, categoryTree := range allCategoryTree {
		remainAllCategory, categoryTree.Children = GetCategoryChildren(remainAllCategory, *categoryTree.Id)
	}

	return remainAllCategory, allCategoryTree
}
