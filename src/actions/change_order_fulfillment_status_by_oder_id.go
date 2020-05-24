package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ChangeOrderFulfilmentStatusByOrderId(orderId int64, body *model.BodyOrder) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.FulfillmentStatus != nil {
		if !model.IsValidFulfillmentStatus(*body.FulfillmentStatus) {

			return false, errors.New("invalid fulfillment_status")
		}
		set = append(set, " fulfillment_status=?")
		args = append(args, body.FulfillmentStatus)
	} else {
		return false, errors.New("require order's fulfillment_status")
	}

	if *body.FulfillmentStatus == "fulfilled" {
		queryString += " fulfilled_at=NOW()"
	}

	if *body.FulfillmentStatus == "cancelled" {
		queryString += " cancelled_at=NOW()"
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	}

	queryString += "WHERE _id=?"
	args = append(args, orderId)

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return false, err
	}

	if rowEffected == nil {
		return false, errors.New("update error")
	}

	return true, nil
}
