package actions

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func UpdateCategory(categoryId int64, body *model.BodyCategory) (*model.Category, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Name != nil {
		set = append(set, " name=?")
		args = append(args, body.Name)
	}

	if body.ImagePath != nil {
		set = append(set, " image_path=?")
		args = append(args, body.ImagePath)
	}

	if body.ParentId != nil {
		if *body.ParentId == 0 {
			set = append(set, " parent_id=?")
			args = append(args, &sql.NullInt64{
				Int64: *body.ParentId,
				Valid: false,
			})
		} else {
			parentCat, err := factories.FindCategoryById(*body.ParentId)

			if err != nil {
				return nil, err
			}

			if parentCat == nil {
				return nil, errors.New("parent category does not exists")
			}

			set = append(set, " parent_id=?")
			args = append(args, &sql.NullInt64{
				Int64: *body.ParentId,
				Valid: true,
			})
		}
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		category, err := factories.FindCategoryById(categoryId)

		if err != nil {
			return nil, err
		}

		if category == nil {
			return nil, errors.New("category does not exists")
		}

		return category, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, categoryId)

	rowEffected, err := factories.UpdateCategory(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetCategoryById(categoryId)
}
