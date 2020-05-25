package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func CheckoutOrder(orderId int64, body *model.BodyCheckoutOrder) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Customer == nil {
		return false, errors.New("customer's information is required")
	}

	if body.Customer.Phone == nil {
		return false, errors.New("customer's phone is required")
	} else {
		set = append(set, " phone=?")
		args = append(args, body.Customer.Phone)
	}

	if body.Customer.Address == nil {
		return false, errors.New("customer's address is required")
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

	queryCustomer := connect.QueryMySQL{
		QueryString: "WHERE phone=?",
		Args:        []interface{}{body.Customer.Phone},
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
		totalItem, totalPrice, err := factories.CountAndCaculateOrderItem(&connect.QueryMySQL{
			QueryString: "WHERE order_id=? AND EXISTS (SELECT _id FROM orders WHERE _id=? AND payment_status='pending')",
			Args:        []interface{}{&orderId, &orderId},
		})

		if err != nil {
			rejectChan <- err
		}

		// log.Println(totalItem)

		if totalItem <= 0 {
			rejectChan <- errors.New("your order is empty or has been checkout")
		}

		resolveChan <- &totalPrice
	}()

	var customer *model.Customer
	var totalPrice *float64

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case *model.Customer:
				customer = val
			case *float64:
				totalPrice = val
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

	setUpdateOrder = append(setUpdateOrder, ` status="payment"`)
	setUpdateOrder = append(setUpdateOrder, ` payment_status="paid"`)
	setUpdateOrder = append(setUpdateOrder, ` paid_at=NOW()`)
	setUpdateOrder = append(setUpdateOrder, ` customer_id=?`)
	argsUpdateOrder = append(argsUpdateOrder, &customerId)
	setUpdateOrder = append(setUpdateOrder, ` total_price=?`)
	argsUpdateOrder = append(argsUpdateOrder, totalPrice)

	if body.Note != nil {
		setUpdateOrder = append(setUpdateOrder, ` note=?`)
		argsUpdateOrder = append(argsUpdateOrder, body.Note)
	}

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
