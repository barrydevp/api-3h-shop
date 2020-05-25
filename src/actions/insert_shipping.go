package actions

import (
	"errors"
	"strings"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertShipping(body *model.BodyShipping) (*model.Shipping, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.OrderId != nil {
		order, err := factories.FindOneOrder(&connect.QueryMySQL{
			QueryString: "WHERE _id=? AND payment_status='pending'",
			Args:        []interface{}{body.OrderId},
		})

		if err != nil {
			return nil, err
		}

		if order == nil {
			return nil, errors.New("order does not exists")
		}
		set = append(set, " order_id=?")
		args = append(args, body.OrderId)
	} else {
		return nil, errors.New("shipping's order_id is required")
	}

	if body.Carrier != nil {
		set = append(set, " carrier=?")
		args = append(args, body.Carrier)
	} else {
		return nil, errors.New("shipping's carrier is required")
	}

	if body.Price != nil {
		set = append(set, " price=?")
		args = append(args, body.Price)
	} else {
		return nil, errors.New("shipping's price is required")
	}

	if body.Note != nil {
		set = append(set, " note=?")
		args = append(args, body.Note)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertShipping(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindShippingById(*id)
}
