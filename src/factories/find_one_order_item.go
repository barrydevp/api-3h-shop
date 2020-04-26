package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOneOrderItem(query *connect.QueryMySQL) (*model.OrderItem, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, product_id, product_item_id, order_id, quantity, status, created_at, updated_at
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

	var _orderItem model.OrderItem

	err = stmt.QueryRow(args...).Scan(
		&_orderItem.RawId,
		&_orderItem.RawProductId,
		&_orderItem.RawProductItemId,
		&_orderItem.RawOrderId,
		&_orderItem.RawQuantity,
		&_orderItem.RawStatus,
		&_orderItem.RawCreatedAt,
		&_orderItem.RawUpdatedAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_orderItem.FillResponse()

		return &_orderItem, nil
	default:
		return nil, err
	}

	return nil, nil
}
