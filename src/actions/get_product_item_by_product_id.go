package actions

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/factories"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetProductItemByProductId(productId int64) ([]*model.ProductItem, error) {
	query := connect.QueryMySQL{
		QueryString: "WHERE product_id=?",
		Args:        []interface{}{productId},
	}

	resolveChan := make(chan interface{}, 2)
	rejectChan := make(chan error)

	go func() {
		data, err := factories.FindProductItem(&query)

		if err != nil {
			rejectChan <- err
		} else {
			resolveChan <- data
		}
	}()

	go func() {
		product, err := factories.FindProductById(productId)

		if err != nil {
			rejectChan <- err
		}

		resolveChan <- product
	}()

	var items []*model.ProductItem
	var product *model.Product

	for i := 0; i < 2; i++ {
		select {
		case data := <-resolveChan:
			switch val := data.(type) {
			case *model.Product:
				product = val
				if product == nil {
					return nil, errors.New("product does not exists")
				}
			case []*model.ProductItem:
				items = val
			}
		case err := <-rejectChan:
			return nil, err
		}
	}

	return items, nil
}
