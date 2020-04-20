package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindProductItem(query *connect.QueryMySQL) ([]*model.ProductItem, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, product_id, stock, in_price, created_at, updated_at, expired_at
		FROM product_items
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString
		args = query.Args
	}

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listProductItem []*model.ProductItem

	for rows.Next() {
		_productItem := model.ProductItem{}

		err = rows.Scan(
			&_productItem.RawId,
			&_productItem.RawProductId,
			&_productItem.RawStock,
			&_productItem.RawInPrice,
			&_productItem.RawCreatedAt,
			&_productItem.RawUpdatedAt,
			&_productItem.RawExpiredAt,
		)

		if err != nil {
			return nil, err
		}

		_productItem.FillResponse()

		listProductItem = append(listProductItem, &_productItem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listProductItem, nil
}
