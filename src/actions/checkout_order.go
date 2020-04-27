package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"log"
	"strings"
)

func CheckoutOrder(orderId int64, body *model.BodyCheckoutOrder) (*model.Order, error) {
	queryString := ""
	var args []interface{}

	var set []string

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
		totalItem, err := factories.CountOrderItem(&connect.QueryMySQL{
			QueryString: "WHERE order_id=? AND EXISTS (SELECT _id FROM orders WHERE _id=order_items.order_id AND status='pending')",
			Args: []interface{}{&orderId},
		})

		if err != nil {
			rejectChan <-err
		}

		log.Println(totalItem)

		if totalItem <= 0 {
			rejectChan <-errors.New("your order is empty or has been checkout")
		}

		resolveChan <-totalItem
	}()


	var customer *model.Customer

	for i := 0; i < 2; i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case *model.Customer:
				customer = val
			case int:
			}

		case err := <-rejectChan:
			return nil, err
		}
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	var customerId int64

	if customer != nil {
		_, err := factories.UpdateCustomer(&connect.QueryMySQL{
			QueryString: queryString + "WHERE _id=?",
			Args:        append(args, customer.Id),
		})

		if err != nil {
			return nil, err
		}

		customerId = *customer.Id
	} else {
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

		customerId = *id
	}

	bodyOrder := model.BodyOrder{}
	status := "payment"
	bodyOrder.Status = &status
	bodyOrder.CustomerId = &customerId
	if body.Note != nil {
		bodyOrder.Note = body.Note
	}

	return UpdateOrder(orderId, &bodyOrder)
}
