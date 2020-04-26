package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func InsertOrderItem(body *model.BodyOrderItem) (*model.OrderItem, error) {
	queryString := ""
	var args []interface{}

	var set []string

	var goroutines []func()
	resolveChan := make(chan interface{}, 4)
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
	} else {
		return nil, errors.New("order_item's product_id is required")
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

		//goroutines = append(goroutines, func() {
		//	orderItem, err := factories.FindOrderItem(&connect.QueryMySQL{
		//		QueryString: "WHERE order_id=? AND product_id=?",
		//		Args: []interface{}{body.OrderId, body.ProductId},
		//	})
		//
		//	if err != nil {
		//		rejectChan <- err
		//		return
		//	}
		//
		//	resolveChan <- orderItem
		//})
	} else {
		return nil, errors.New("order_item's product_id is required")
	}

	if body.Quantity != nil {
		set = append(set, " quantity=?")
		args = append(args, body.Quantity)
	} else {
		return nil, errors.New("order_item's quantity is required")
	}

	for _, goroutine := range goroutines {
		go goroutine()
	}

	for i := 0; i < len(goroutines); i++ {
		select {
		case res := <-resolveChan:
			switch val := res.(type) {
			case *model.OrderItem:
				if val != nil {
					_, err := factories.UpdateOrderItem(&connect.QueryMySQL{
						QueryString: "SET quantity=quantity + ? WHERE _id=?",
						Args:        []interface{}{body.Quantity, val.Id},
					})

					if err != nil {
						return nil, err
					}

					return factories.FindOrderItemById(*val.Id)
				}
			default:
			}
		case err := <-rejectChan:
			return nil, err
		}
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertOrderItem(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindOrderItemById(*id)
}
