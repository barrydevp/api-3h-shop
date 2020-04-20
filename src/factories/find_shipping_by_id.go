package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindShippingById(shippingId int64) (*model.Shipping, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, carrier, status, order_id, created_at, updated_at, delivered_at
		FROM shippings
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _shipping model.Shipping

	err = stmt.QueryRow(shippingId).Scan(
		&_shipping.RawId,
		&_shipping.RawCarrier,
		&_shipping.RawStatus,
		&_shipping.RawStatus,
		&_shipping.RawCreatedAt,
		&_shipping.RawUpdatedAt,
		&_shipping.RawDeliveredAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_shipping.FillResponse()

		return &_shipping, nil
	default:
		return nil, err
	}

	return nil, nil
}
