package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOrderById(orderId int64) (*model.Order, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, session, customer_id, status, payment_status, fulfillment_status, created_at, updated_at, paid_at, fulfilled_at, cancelled_at, note
		FROM orders
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _order model.Order

	err = stmt.QueryRow(orderId).Scan(
		&_order.RawId,
		&_order.RawSession,
		&_order.RawCustomerId,
		&_order.RawPaymentStatus,
		&_order.RawStatus,
		&_order.RawFulfillmentStatus,
		&_order.RawCreatedAt,
		&_order.RawUpdatedAt,
		&_order.RawPaidAt,
		&_order.RawFulfilledAt,
		&_order.RawCancelledAt,
		&_order.RawNote,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_order.FillResponse()

		return &_order, nil
	default:
		return nil, err
	}

	return nil, nil
}
