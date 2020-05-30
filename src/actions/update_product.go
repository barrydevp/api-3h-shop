package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateProduct(productId int64, body *model.BodyProduct) (*model.Product, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.CategoryId != nil {
		//category, err := factories.FindCategoryById(*body.CategoryId)
		//
		//if err != nil {
		//	return nil, err
		//}
		//
		//if category == nil {
		//	return nil, errors.New("category does not exists")
		//}

		set = append(set, " category_id=?")
		args = append(args, body.CategoryId)
	}

	if body.Name != nil {
		set = append(set, " name=?")
		args = append(args, body.Name)
	}

	if body.OutPrice != nil {
		set = append(set, " out_price=?")
		args = append(args, body.OutPrice)
	}

	if body.Discount != nil {
		set = append(set, " discount=?")
		args = append(args, body.Discount)
	}

	if body.ImagePath != nil {
		set = append(set, " image_path=?")
		args = append(args, body.ImagePath)
	}

	if body.Description != nil {
		set = append(set, " description=?")
		args = append(args, body.Description)
	}

	if body.Tags != nil {
		set = append(set, " tags=tags+','+?")
		args = append(args, body.Tags)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		product, err := factories.FindProductById(productId)

		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, errors.New("product does not exists")
		}

		return product, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, productId)

	rowEffected, err := factories.UpdateProduct(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetProductById(productId)
}
