package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCustomer(query *connect.QueryMySQL) ([]*model.Customer, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, phone, address, full_name, email, updated_at
		FROM customers
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
	var listCustomer []*model.Customer

	for rows.Next() {
		_customer := model.Customer{}

		err = rows.Scan(
			&_customer.RawId,
			&_customer.RawPhone,
			&_customer.RawAddress,
			&_customer.RawFullName,
			&_customer.RawEmail,
			&_customer.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_customer.FillResponse()

		listCustomer = append(listCustomer, &_customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listCustomer, nil
}
