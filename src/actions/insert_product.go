package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertProduct(body *model.BodyProduct) (*model.Product, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.CategoryId == nil {
		return nil, errors.New("product's category_id is required")
	} else {
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

	if body.Name == nil {
		return nil, errors.New("product's name is required")
	} else {
		set = append(set, " name=?")
		args = append(args, body.Name)
	}

	if body.OutPrice == nil {
		return nil, errors.New("product's out_price is required")
	} else {
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

	tagsFromName := strings.Split(*body.Name, " ")

	for i := 0; i < len(tagsFromName); i++ {
		current := tagsFromName[i]
		for j := i + 1; j < len(tagsFromName); j++ {
			current += tagsFromName[j]
			tagsFromName = append(tagsFromName, current)
		}
	}

	set = append(set, " tags=?")

	if body.Tags != nil {
		tagsFromName = append(tagsFromName, *body.Tags)
	}

	args = append(args, strings.Join(tagsFromName, ","))

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertProduct(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindProductById(*id)
}
