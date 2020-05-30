package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func StatisticOrder(query *connect.QueryMySQL) (*model.StatisticOrder, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			COUNT(*), SUM(IF(payment_status='pending', 1, 0)), SUM(IF(payment_status='paid', 1, 0)), SUM(IF(fulfillment_status='in-production', 1, 0)), SUM(IF(fulfillment_status='shipped', 1, 0)), SUM(IF(fulfillment_status='cancelled', 1, 0)), SUM(IF(fulfillment_status='fulfilled', 1, 0))
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

	var statistic model.StatisticOrder

	err = stmt.QueryRow(args...).Scan(
		&statistic.TotalOrder,
		&statistic.Pending,
		&statistic.Paid,
		&statistic.InProduction,
		&statistic.Shipped,
		&statistic.Cancelled,
		&statistic.Fulfilled,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &statistic, nil
	default:
		return nil, err
	}

	return &statistic, nil
}
