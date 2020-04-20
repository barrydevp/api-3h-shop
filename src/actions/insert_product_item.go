package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func InsertProductItem(body *model.BodyProductItem) (*model.ProductItem, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.ProductId == nil {
		return nil, errors.New("product_item's product_id is required")
	} else {
		product, err := factories.FindProductById(*body.ProductId)

		if err != nil {
			return nil, err
		}

		if product == nil {
			return nil, errors.New("product does not exists")
		}
		set = append(set, " product_id=?")
		args = append(args, body.ProductId)
	}

	if body.Stock == nil {
		return nil, errors.New("product_item's stock is required")
	} else {
		set = append(set, " stock=?")
		args = append(args, body.Stock)
	}

	if body.InPrice == nil {
		return nil, errors.New("product_item's in_price is required")
	} else {
		set = append(set, " in_price=?")
		args = append(args, body.InPrice)
	}

	if body.ExpiredAt == nil {
		return nil, errors.New("product_item's expired_at is required")
	} else {
		set = append(set, " expired_at=?")
		args = append(args, body.ExpiredAt)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + ", created_at=NOW() \n"
	} else {
		return nil, errors.New("invalid body")
	}

	id, err := factories.InsertProductItem(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, errors.New("insert error")
	}

	return factories.FindProductItemById(*id)
}
