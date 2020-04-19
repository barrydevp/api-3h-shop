package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func InsertCustomer(body *model.BodyCustomer) (*model.Customer, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Phone == nil {
		return nil, errors.New("customer's phone is required")
	} else {
		set = append(set, " phone=?")
		args = append(args, body.Phone)
	}

	if body.Address == nil {
		return nil, errors.New("customer's address is required")
	} else {
		set = append(set, " address=?")
		args = append(args, body.Address)
	}

	if body.FullName != nil {
		set = append(set, " full_name=?")
		args = append(args, body.FullName)
	}

	if body.Email != nil {
		set = append(set, " email=?")
		args = append(args, body.Email)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertCustomer(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindCustomerById(*id)
}
