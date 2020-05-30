package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindWarranty(query *connect.QueryMySQL) ([]*model.Warranty, error) {
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

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listWarranty []*model.Warranty

	for rows.Next() {
		_warranty := model.Warranty{}

		err = rows.Scan(
			&_warranty.RawId,
			&_warranty.RawCode,
			&_warranty.RawMonth,
			&_warranty.RawTrial,
			&_warranty.RawStatus,
			&_warranty.RawDescription,
			&_warranty.RawCategoryId,
		)

		if err != nil {
			return nil, err
		}

		_warranty.FillResponse()

		listWarranty = append(listWarranty, &_warranty)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listWarranty, nil
}
