package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetOrderAndTotalItemWithSession(query *connect.QueryMySQL) (map[string]int, map[string]*model.Order, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, session, customer_id, status, total_price, payment_status, fulfillment_status, created_at, updated_at, paid_at, fulfilled_at, cancelled_at, note,
			(	
				SELECT COUNT(*) 
				FROM order_items
				WHERE order_id = orders._id
			) 
		FROM orders
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString
		args = query.Args
	}

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return nil, nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()
	mapTotal := make(map[string]int)
	mapOrders := make(map[string]*model.Order)

	for rows.Next() {
		_order := model.Order{}
		total := 0

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
			&total,
		)

		if err != nil {
			return nil, nil, err
		}

		_order.FillResponse()

		mapTotal[*_order.Session] = total
		mapOrders[*_order.Session] = &_order
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return mapTotal, mapOrders, nil
}
