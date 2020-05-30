package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
)

func MarkOrderPaid(orderId int64) (bool, error) {
	queryString := ""

	totalItem, totalPrice, err := factories.CountAndCaculateOrderItem(&connect.QueryMySQL{
		QueryString: "WHERE order_id=? AND EXISTS (SELECT _id FROM orders WHERE _id=order_items.order_id AND payment_status='pending')",
		Args:        []interface{}{&orderId},
	})

	if err != nil {
		return false, err
	}

	// log.Println(totalItem)

	if totalItem <= 0 {
		return false, errors.New("your order is empty or has been checkout")
	}

	var argsUpdateOrder []interface{}
	var setUpdateOrder []string

	setUpdateOrder = append(setUpdateOrder, ` status="mark_paid"`)
	setUpdateOrder = append(setUpdateOrder, ` payment_status="paid"`)
	setUpdateOrder = append(setUpdateOrder, ` paid_at=NOW()`)
	setUpdateOrder = append(setUpdateOrder, ` total_price=?`)
	argsUpdateOrder = append(argsUpdateOrder, totalPrice)

	if len(setUpdateOrder) > 0 {
		queryString = "SET" + strings.Join(setUpdateOrder, ",") + "\n"
	} else {
		return false, errors.New("invalid body order")
	}

	queryString += "WHERE _id=?"
	argsUpdateOrder = append(argsUpdateOrder, &orderId)

	// log.Println(queryString)

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
