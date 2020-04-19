package factories

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindAllCategoryTree() ([]*model.CategoryTree, error) {
	allCategory, err := FindCategory(nil)

	if err != nil {
		return nil, err
	}

	tree := MakeCategoryTree(allCategory)

	return tree, nil
}

func MakeCategoryTree(all []*model.Category) []*model.CategoryTree {
	var allCategoryTree []*model.CategoryTree
	var remainAllCategory []*model.Category

	for _, category := range all {
		if category.ParentId == nil {
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

	return allCategoryTree
}
