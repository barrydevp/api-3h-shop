package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetOrderCustomerByOrderId(orderId int64) (*model.Customer, error) {
	query := connect.QueryMySQL{
		QueryString: "WHERE _id=(SELECT customer_id FROM orders WHERE _id=?)",
		Args:        []interface{}{orderId},
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		customer, err := factories.FindOneCustomer(&query)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- customer
		}
	}()

	go func() {
		order, err := factories.FindOrderById(orderId)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- order
	}()

	var customer *model.Customer
	var order *model.Order

	for i := 0; i < 2; i++ {
		select {
		case data := <-resolveChan:
			switch val := data.(type) {
			case *model.Order:
				order = val
				if order == nil {
					return nil, errors.New("order does not exists")
				}
			case *model.Customer:
				customer = val
			}
		case err := <-rejectChan:
			return nil, err
		}
	}

	return customer, nil
}
