package actions

import (
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertProductItemByProductId(productId int64, body *model.BodyProductItem) (*model.ProductItem, error) {
	body.ProductId = &productId

	return InsertProductItem(body)
}
