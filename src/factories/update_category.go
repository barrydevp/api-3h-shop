package factories

import (
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
)

func UpdateCategory(query *connect.QueryMySQL) (*int64, error) {
	if query == nil {
		return nil, errors.New("query is required")
	}

	connection := connections.Mysql.GetConnection()

	queryString := `
		UPDATE categories 
	` + query.QueryString
	args := query.Args

	stmt, err := connection.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &id, nil
}
