package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func UpdateCoupon(couponId int64, body *model.BodyCoupon) (*model.Coupon, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Description != nil {
		set = append(set, " description=?")
		args = append(args, body.Description)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		coupon, err := factories.FindCouponById(couponId)

		if err != nil {
			return nil, err
		}

		if coupon == nil {
			return nil, errors.New("coupon does not exists")
		}

		return coupon, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, couponId)

	rowEffected, err := factories.UpdateCoupon(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetCouponById(couponId)
}
