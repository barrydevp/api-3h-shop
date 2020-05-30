package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOneOrder(query *connect.QueryMySQL) (*model.Order, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, session, customer_id, status, total_price, payment_status, fulfillment_status, created_at, updated_at, paid_at, fulfilled_at, cancelled_at, note, coupon_id
		FROM orders
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

	var _order model.Order

	err = stmt.QueryRow(args...).Scan(
		&_order.RawId,
		&_order.RawSession,
		&_order.RawCustomerId,
		&_order.RawStatus,
		&_order.RawTotalPrice,
		&_order.RawPaymentStatus,
		&_order.RawFulfillmentStatus,
		&_order.RawCreatedAt,
		&_order.RawUpdatedAt,
		&_order.RawPaidAt,
		&_order.RawFulfilledAt,
		&_order.RawCancelledAt,
		&_order.RawNote,
		&_order.RawCouponId,
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
