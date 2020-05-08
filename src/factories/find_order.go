package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOrder(query *connect.QueryMySQL) ([]*model.Order, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, session, customer_id, status, total_price, payment_status, fulfillment_status, created_at, updated_at, paid_at, fulfilled_at, cancelled_at, note
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

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listOrder []*model.Order

	for rows.Next() {
		_order := model.Order{}

		err = rows.Scan(
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
		)

		if err != nil {
			return nil, err
		}

		_order.FillResponse()

		listOrder = append(listOrder, &_order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listOrder, nil
}
