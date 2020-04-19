package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCustomerById(customerId int64) (*model.Customer, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, phone, address, full_name, email, updated_at
		FROM customers
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _customer model.Customer

	err = stmt.QueryRow(customerId).Scan(
		&_customer.RawId,
		&_customer.RawPhone,
		&_customer.RawAddress,
		&_customer.RawFullName,
		&_customer.RawEmail,
		&_customer.RawUpdatedAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_customer.FillResponse()

		return &_customer, nil
	default:
		return nil, err
	}

	return nil, nil
}
