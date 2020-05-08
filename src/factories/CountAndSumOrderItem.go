package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
)

func CountAndCaculateOrderItem(query *connect.QueryMySQL) (int, float64, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			COUNT(*), ROUND(SUM(products.out_price * (1 - products.discount)), 2)
		FROM order_items
		INNER JOIN products
		ON order_items.product_id = products._id
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString
		args = query.Args
	}

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return 0, 0, err
	}

	defer stmt.Close()

	var count int
	var totalPrice float64

	err = stmt.QueryRow(args...).Scan(&count, &totalPrice)

	switch err {
	case sql.ErrNoRows:
		return 0, 0, nil
	case nil:
		return count, totalPrice, nil
	default:
		return 0, 0, err
	}

	return count, totalPrice, nil
}
