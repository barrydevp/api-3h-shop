package factories

import (
	"database/sql"
	"log"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
)

func CountAndCaculateOrderItem(query *connect.QueryMySQL) (int, float64, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT o.count, ROUND(o.total_price * (100 - COALESCE((SELECT discount FROM coupons WHERE _id=(SELECT coupon_id FROM orders WHERE _id=?)), 0)) / 100, 2)
		FROM (
			SELECT
				COUNT(*) count, COALESCE(SUM(products.out_price * o_i.quantity * (1 - products.discount)), 0) total_price
`
	var args []interface{}

	if query != nil {
		queryString += "FROM (SELECT * FROM order_items " + query.QueryString + ") o_i"
		args = query.Args
	}
	queryString += `
			INNER JOIN products
			ON o_i.product_id = products._id
		) o
	`

	log.Println(queryString)

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
