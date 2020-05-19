package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
)

func CountOrderItemTwoSession(query *connect.QueryMySQL) ([]int, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			COUNT(*)
		FROM order_items
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString
		args = query.Args
	}

	queryString += `GROUP BY order_id`

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
	var listTotal []int

	for rows.Next() {
		total := 0

		err = rows.Scan(&total)

		if err != nil {
			return nil, err
		}

		listTotal = append(listTotal, total)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listTotal, nil
}
