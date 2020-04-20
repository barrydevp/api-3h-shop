package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindProductItemById(productItemId int64) (*model.ProductItem, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, product_id, stock, in_price, created_at, updated_at, expired_at
		FROM product_items
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _productItem model.ProductItem

	err = stmt.QueryRow(productItemId).Scan(
		&_productItem.RawId,
		&_productItem.RawProductId,
		&_productItem.RawStock,
		&_productItem.RawInPrice,
		&_productItem.RawCreatedAt,
		&_productItem.RawUpdatedAt,
		&_productItem.RawExpiredAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_productItem.FillResponse()

		return &_productItem, nil
	default:
		return nil, err
	}

	return nil, nil
}
