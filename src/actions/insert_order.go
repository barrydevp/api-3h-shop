package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertOrder(body *model.BodyOrder) (*model.Order, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.Session != nil {
		set = append(set, " session=?")
		args = append(args, body.Session)
	} else {
		return nil, errors.New("order's session is required")
	}

	if body.Note != nil {
		set = append(set, " note=?")
		args = append(args, body.Note)
	}

	if body.CouponId != nil {
		set = append(set, " coupon_id=?")
		args = append(args, body.CouponId)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertOrder(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindOrderById(*id)
}
