package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func InsertShipping(body *model.BodyShipping) (*model.Shipping, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.OrderId != nil {
		order, err := factories.FindOrderById(*body.OrderId)

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
