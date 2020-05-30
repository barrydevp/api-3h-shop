package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOrderItem(query *connect.QueryMySQL) ([]*model.OrderItem, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, product_id, product_item_id, order_id, quantity, status, created_at, updated_at, warranty_id
		FROM order_items
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
	var listOrderItem []*model.OrderItem

	for rows.Next() {
		_orderItem := model.OrderItem{}

		err = rows.Scan(
			&_orderItem.RawId,
			&_orderItem.RawProductId,
			&_orderItem.RawProductItemId,
			&_orderItem.RawOrderId,
			&_orderItem.RawQuantity,
			&_orderItem.RawStatus,
			&_orderItem.RawCreatedAt,
			&_orderItem.RawUpdatedAt,
			&_orderItem.RawWarrantyId,
		)

		if err != nil {
			return nil, err
		}

		_orderItem.FillResponse()

		listOrderItem = append(listOrderItem, &_orderItem)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listOrderItem, nil
}
