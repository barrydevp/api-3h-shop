package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func UpdateOrderItem(orderItemId int64, body *model.BodyOrderItem) (*model.OrderItem, error) {
	queryString := ""
	var args []interface{}

	var set []string

	var goroutines []func()
	resolveChan := make(chan interface{}, 3)
	rejectChan := make(chan error)

	if body.ProductId != nil {
		//goroutines = append(goroutines, func() {
		//	order, err := factories.FindProductById(*body.ProductId)
		//
		//	if err != nil {
		//		rejectChan <- err
		//		return
		//	}
		//	if order == nil {
		//		rejectChan <- errors.New("product does not exists")
		//		return
		//	}
		//
		//	resolveChan <- order
		//})

		set = append(set, " product_id=?")
		args = append(args, body.ProductId)
	}

	if body.ProductItemId != nil {
		//goroutines = append(goroutines, func() {
		//	order, err := factories.FindProductItemById(*body.ProductItemId)
		//
		//	if err != nil {
		//		rejectChan <- err
		//		return
		//	}
		//	if order == nil {
		//		rejectChan <- errors.New("product_item does not exists")
		//		return
		//	}
		//
		//	resolveChan <- order
		//})

		set = append(set, " product_item_id=?")
		args = append(args, body.ProductItemId)
	}

	if body.OrderId != nil {
		//goroutines = append(goroutines, func() {
		//	order, err := factories.FindOrderById(*body.OrderId)
		//
		//	if err != nil {
		//		rejectChan <- err
		//		return
		//	}
		//	if order == nil {
		//		rejectChan <- errors.New("order does not exists")
		//		return
		//	}
		//
		//	resolveChan <- order
		//})

		set = append(set, " order_id=?")
		args = append(args, body.OrderId)
	}

	for _, goroutine := range goroutines {
		go goroutine()
	}

	for i := 0; i < len(goroutines); i++ {
		select {
		case <-resolveChan:
		case err := <-rejectChan:
			return nil, err
		}
	}

	if body.Quantity != nil {
		set = append(set, " quantity=?")
		args = append(args, body.Quantity)
	}

	if body.Status != nil {
		set = append(set, " status=?")
		args = append(args, body.Quantity)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		orderItem, err := factories.FindOrderItemById(orderItemId)

		if err != nil {
			return nil, err
		}

		if orderItem == nil {
			return nil, errors.New("orderItem does not exists")
		}

		return orderItem, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, orderItemId)

	rowEffected, err := factories.UpdateOrderItem(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetOrderItemById(orderItemId)
}
