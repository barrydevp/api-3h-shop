package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOneWarranty(query *connect.QueryMySQL) (*model.Warranty, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, code, month, trial, status, description, category_id
		FROM warranties
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

	var _warranty model.Warranty

	err = stmt.QueryRow(args...).Scan(
		&_warranty.RawId,
		&_warranty.RawCode,
		&_warranty.RawMonth,
		&_warranty.RawTrial,
		&_warranty.RawStatus,
		&_warranty.RawDescription,
		&_warranty.RawCategoryId,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_warranty.FillResponse()

		return &_warranty, nil
	default:
		return nil, err
	}

	return nil, nil

}
