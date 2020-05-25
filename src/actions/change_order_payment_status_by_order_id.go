package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func ChangeOrderPaymentStatusByOrderId(orderId int64, body *model.BodyOrder) (bool, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.PaymentStatus != nil {
		if !model.IsValidPaymentStatus(*body.PaymentStatus) {
			return false, errors.New("invalid payment_status")
		}
		set = append(set, " payment_status=?")
		args = append(args, body.PaymentStatus)
	} else {
		return false, errors.New("require order's payment_status")
	}

	if *body.PaymentStatus == "paid" {
		set = append(set, " paid_at=NOW()")
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
