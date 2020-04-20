package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func UpdateShipping(shippingId int64, body *model.BodyShipping) (*model.Shipping, error) {
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
	}

	if body.Carrier != nil {
		set = append(set, " carrier=?")
		args = append(args, body.Carrier)
	}

	if body.DeliveredAt != nil {
		set = append(set, " delivered_at=?")
		args = append(args, body.DeliveredAt)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		shipping, err := factories.FindShippingById(shippingId)

		if err != nil {
			return nil, err
		}

		if shipping == nil {
			return nil, errors.New("shipping does not exists")
		}

		return shipping, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, shippingId)

	rowEffected, err := factories.UpdateShipping(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetShippingById(shippingId)
}
