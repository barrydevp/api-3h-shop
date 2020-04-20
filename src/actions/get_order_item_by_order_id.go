package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetOrderItemByOrderId(orderId int64) ([]*model.OrderItem, error) {
	query := connect.QueryMySQL{
		QueryString: "WHERE order_id=?",
		Args:        []interface{}{orderId},
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindOrderItem(&query)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		order, err := factories.FindOrderById(orderId)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- order
	}()

	var items []*model.OrderItem
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
			case []*model.OrderItem:
				items = val
			}
		case err := <-rejectChan:
			return nil, err
		}
	}

	return items, nil
}
