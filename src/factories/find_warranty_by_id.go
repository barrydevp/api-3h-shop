package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindWarrantyById(warrantyId int64) (*model.Warranty, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, code, month, trial, status, description, category_id
		FROM warranties
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _warranty model.Warranty

	err = stmt.QueryRow(warrantyId).Scan(
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
