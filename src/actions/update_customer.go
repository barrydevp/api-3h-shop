package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func UpdateCustomer(customerId int64, body *model.BodyCustomer) (*model.Customer, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Phone != nil  {
		set = append(set, " phone=?")
		args = append(args, body.Phone)
	}

	if body.Address != nil {
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
		customer, err := factories.FindCustomerById(customerId)

		if err != nil {
			return nil, err
		}

		if customer == nil {
			return nil, errors.New("customer does not exists")
		}

		return customer, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, customerId)

	rowEffected, err := factories.UpdateCustomer(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetCustomerById(customerId)
}
