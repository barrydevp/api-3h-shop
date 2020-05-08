package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateOrder(orderId int64, body *model.BodyOrder) (*model.Order, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.CustomerId != nil {
		set = append(set, " customer_id=?")
		args = append(args, body.CustomerId)
	}

	if body.Status != nil {
		set = append(set, " status=?")
		args = append(args, body.Status)
	}

	if body.TotalPrice != nil {
		set = append(set, " total_price=?")
		args = append(args, body.TotalPrice)
	}

	if body.PaymentStatus != nil {
		set = append(set, " payment_status=?")
		args = append(args, body.PaymentStatus)
	}

	if body.FulfillmentStatus != nil {
		set = append(set, " fulfilment_status=?")
		args = append(args, body.FulfillmentStatus)
	}

	if body.PaidAt != nil {
		set = append(set, " paid_at=?")
		args = append(args, body.PaidAt)
	}

	if body.Note != nil {
		set = append(set, " note=?")
		args = append(args, body.Note)
	}

	if body.FulfilledAt != nil {
		set = append(set, " fulfilled_at=?")
		args = append(args, body.FulfilledAt)
	}

	if body.CancelledAt != nil {
		set = append(set, " cancelled_at=?")
		args = append(args, body.CancelledAt)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		order, err := factories.FindOrderById(orderId)

		if err != nil {
			return nil, err
		}

		if order == nil {
			return nil, errors.New("order does not exists")
		}

		return order, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, orderId)

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetOrderById(orderId)
}
