package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetOrderCoupon(orderId int64) (*model.Coupon, error) {
	order, err := factories.FindOrderById(orderId)

	if order == nil {
		return nil, errors.New("order does not exists")
	}

	if order.CouponId == nil {
		return nil, nil
	}

	coupon, err := factories.FindCouponById(*order.CouponId)

	if err != nil {
		return nil, err
	}

	return coupon, nil
}
