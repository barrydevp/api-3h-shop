package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOrderItemById(orderItemId int64) (*model.OrderItem, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, product_id, product_item_id, order_id, quantity, status, created_at, updated_at, warranty_id
		FROM order_items
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _orderItem model.OrderItem

	err = stmt.QueryRow(orderItemId).Scan(
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
