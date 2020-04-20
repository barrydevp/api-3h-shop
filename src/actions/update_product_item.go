package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
	"strings"
)

func UpdateProductItem(productItemId int64, body *model.BodyProductItem) (*model.ProductItem, error) {
	queryString := ""
	var args []interface{}

	var set []string

	if body.ProductId != nil {
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

	if body.Stock != nil {
		set = append(set, " stock=?")
		args = append(args, body.Stock)
	}

	if body.InPrice != nil {
		set = append(set, " in_price=?")
		args = append(args, body.InPrice)
	}

	if body.ExpiredAt != nil {
		set = append(set, " expired_at=?")
		args = append(args, body.ExpiredAt)
	}

	if len(set) > 0 {
		queryString += "SET" + strings.Join(set, ",") + "\n"
	} else {
		productItem, err := factories.FindProductItemById(productItemId)

		if err != nil {
			return nil, err
		}

		if productItem == nil {
			return nil, errors.New("productItem does not exists")
		}

		return productItem, nil
	}

	queryString += "WHERE _id=?"
	args = append(args, productItemId)

	rowEffected, err := factories.UpdateProductItem(&connect.QueryMySQL{
		QueryString: queryString,
		Args:        args,
	})

	if err != nil {
		return nil, err
	}

	if rowEffected == nil {
		return nil, errors.New("update error")
	}

	return GetProductItemById(productItemId)
}
