package actions

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func InsertCategory(body *model.BodyCategory) (*model.Category, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Name == nil {
		return nil, errors.New("category's name is required")
	} else {
		set = append(set, " name=?")
		args = append(args, body.Name)
	}

	if body.ParentId != nil {
		if *body.ParentId == -1 {
			set = append(set, " parent_id=?")
			args = append(args, &sql.NullInt64{
				Int64: *body.ParentId,
				Valid: false,
			})
		} else {
			//parentCat, err := factories.FindCategoryById(*body.ParentId)
			//
			//if err != nil {
			//	return nil, err
			//}
			//
			//if parentCat == nil {
			//	return nil, errors.New("parent category does not exists")
			//}

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
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertCategory(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindCategoryById(*id)
}
