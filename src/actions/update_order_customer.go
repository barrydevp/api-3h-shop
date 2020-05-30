package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateOrderCustomer(orderId int64, body *model.BodyCustomer) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Phone == nil {
		return false, errors.New("customer's phone is required")
	} else {
		set = append(set, " phone=?")
		args = append(args, body.Phone)
	}

	if body.Address == nil {
		return false, errors.New("customer's address is required")
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

	queryCustomer := connect.QueryMySQL{
		QueryString: "WHERE phone=?",
		Args:        []interface{}{body.Phone},
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		existCustomer, err := factories.FindOneCustomer(&queryCustomer)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- existCustomer
		}
	}()

	go func() {
		order, err := factories.FindOrderById(orderId)

		if err != nil {
			rejectChan <- err

			return
		}

		// log.Println(totalItem)

		if order == nil {
			rejectChan <- errors.New("order does not exists")

			return
		}

		resolveChan <- &order
	}()

	var customer *model.Customer

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case *model.Customer:
				customer = val
			case *model.Order:
			}

		case err := <-rejectChan:
			return false, err
		}
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return false, errors.New("invalid body customer")
	}

	var customerId int64

	if customer != nil {
		_, err := factories.UpdateCustomer(&connect.QueryMySQL{
			QueryString: queryString + "WHERE _id=?",
			Args:        append(args, customer.Id),
		})

		if err != nil {
			return false, err
		}

		customerId = *customer.Id
	} else {
		id, err := factories.InsertCustomer(&connect.QueryMySQL{
			QueryString: queryString,
			Args:        args,
		})

		if err != nil {
			return false, err
		}

		if id == nil {
			return false, errors.New("insert error")
		}

		customerId = *id
	}

	var argsUpdateOrder []interface{}
	var setUpdateOrder []string

	setUpdateOrder = append(setUpdateOrder, ` status="update_customer"`)
	setUpdateOrder = append(setUpdateOrder, ` customer_id=?`)
	argsUpdateOrder = append(argsUpdateOrder, &customerId)

	if len(setUpdateOrder) > 0 {
		queryString = "SET" + strings.Join(setUpdateOrder, ",") + "\n"
	} else {
		return false, errors.New("invalid body order")
	}

	queryString += "WHERE _id=?"
	argsUpdateOrder = append(argsUpdateOrder, &orderId)

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        argsUpdateOrder,
	})

	if err != nil {
		return false, err
	}

	if rowEffected == nil {
		return false, errors.New("update error")
	}

	return true, nil
}
