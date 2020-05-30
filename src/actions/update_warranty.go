package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateWarranty(warrantyId int64, body *model.BodyWarranty) (*model.Warranty, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Month != nil {
		set = append(set, " month=?")
		args = append(args, body.Month)
	}

	if body.Trial != nil {
		set = append(set, " trail=?")
		args = append(args, body.Trial)
	}

	if body.Description != nil {
		set = append(set, " description=?")
		args = append(args, body.Description)
	}

	if body.CategoryId != nil {
		set = append(set, " category_id=?")
		args = append(args, body.CategoryId)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		warranty, err := factories.FindWarrantyById(warrantyId)

		if err != nil {
			return nil, err
		}

		if warranty == nil {
			return nil, errors.New("warranty does not exists")
		}

		return warranty, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, warrantyId)

	rowEffected, err := factories.UpdateWarranty(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetWarrantyById(warrantyId)
}
