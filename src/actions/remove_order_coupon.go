package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
)

func RemoveOrderCoupon(orderId int64) (bool, error) {
	existOrder, err := factories.FindOneOrder(&connect.QueryMySQL{
		QueryString: "WHERE _id=? AND payment_status='pending'",
		Args:        []interface{}{&orderId},
	})

	if err != nil {
		return false, err
	}

	if existOrder == nil {
		return false, errors.New("order does not exists or has been checkout")
	}

	rowEffected, err := factories.UpdateOrder(&connect.QueryMySQL{
		QueryString: `SET coupon_id=NULL, status="remove_coupon"`,
		Args:        []interface{}{},
	})

	if err != nil {
		return false, err
	}

	if rowEffected == nil {
		return false, errors.New("update error")
	}

	return true, nil
}
