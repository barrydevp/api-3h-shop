package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertCoupon(body *model.BodyCoupon) (*model.Coupon, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Code == nil {
		return nil, errors.New("coupon's code is required")
	} else {
		set = append(set, " code=?")
		args = append(args, body.Code)
	}

	if body.Discount == nil {
		return nil, errors.New("coupon's discount is required")
	} else {
		set = append(set, " discount=?")
		args = append(args, body.Discount)
	}

	if body.Description != nil {
		set = append(set, " description=?")
		args = append(args, body.Description)
	} else {
		return nil, errors.New("coupon's description is required")
	}

	if body.ExpiresAt != nil {
		set = append(set, " expires_at=?")
		args = append(args, body.ExpiresAt)
	} else {
		return nil, errors.New("coupon's expires_at is required")
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertCoupon(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindCouponById(*id)
}
