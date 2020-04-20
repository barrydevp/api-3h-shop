package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
)

func CountProductItem(query *connect.QueryMySQL) (int, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			COUNT(*)
		FROM product_items
	`
	var args []interface{}

	if query != nil {
		queryString += query.QueryString
		args = query.Args
	}

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var total int

	err = stmt.QueryRow(args...).Scan(&total)

	switch err {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		return total, nil
	default:
		return 0, err
	}

	return total, nil
}
