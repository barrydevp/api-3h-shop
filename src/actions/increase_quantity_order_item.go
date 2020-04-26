package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func IncreaseQuantityOrderItem(orderItemId int64, body *model.BodyOrderItem) (*model.OrderItem, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Quantity != nil {
		set = append(set, " quantity=quantity + ?")
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
