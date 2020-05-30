package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertWarranty(body *model.BodyWarranty) (*model.Warranty, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Code == nil {
		return nil, errors.New("warranty's code is required")
	} else {
		set = append(set, " code=?")
		args = append(args, body.Code)
	}

	if body.Month == nil {
		return nil, errors.New("warranty's month is required")
	} else {
		set = append(set, " month=?")
		args = append(args, body.Month)
	}

	if body.Trial != nil {
		set = append(set, " trial=?")
		args = append(args, body.Trial)
	} else {
		return nil, errors.New("warranty's trail is required")
	}

	if body.Description != nil {
		set = append(set, " description=?")
		args = append(args, body.Description)
	} else {
		return nil, errors.New("warranty's description is required")
	}

	if body.CategoryId != nil {
		set = append(set, " category_id=?")
		args = append(args, body.CategoryId)
	} else {
		return nil, errors.New("warranty's category_id is required")
	}

	set = append(set, " status='active'")

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertWarranty(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindWarrantyById(*id)
}
