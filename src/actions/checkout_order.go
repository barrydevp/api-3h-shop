package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func CheckoutOrder(orderId int64, body *model.BodyCheckoutOrder) (*model.Order, error) {
	queryString := ""
	var args []interface{}

	var set []string

	totalItem, err := factories.CountCategory(&connect.QueryMySQL{
		QueryString: "WHERE order_id=?",
		Args: []interface{}{&orderId},
	})

	if err != nil {
		return nil, err
	}

	if totalItem <= 0 {
		return nil, errors.New("your order is empty")
	}

	if body.Customer == nil {
		return nil, errors.New("customer's information is required")
	}

	if body.Customer.Phone == nil {
		return nil, errors.New("customer's phone is required")
	} else {
		set = append(set, " phone=?")
		args = append(args, body.Customer.Phone)
	}

	if body.Customer.Address == nil {
		return nil, errors.New("customer's address is required")
	} else {
		set = append(set, " address=?")
		args = append(args, body.Customer.Address)
	}

	if body.Customer.FullName != nil {
		set = append(set, " full_name=?")
		args = append(args, body.Customer.FullName)
	}

	if body.Customer.Email != nil {
		set = append(set, " email=?")
		args = append(args, body.Customer.Email)
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

	bodyOrder := model.BodyOrder{}
	status := "payment"
	bodyOrder.Status = &status
	if body.Note != nil {
		bodyOrder.Note = body.Note
	}

	return UpdateOrder(orderId, &bodyOrder)
}
