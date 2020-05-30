package actions

import (
	"errors"

	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetCouponById(couponId int64) (*model.Coupon, error) {
	coupon, err := factories.FindCouponById(couponId)

	if err != nil {
		return nil, err
	}

	if coupon == nil {
		return nil, errors.New("coupon does not exists")
	}

	return coupon, nil
}
