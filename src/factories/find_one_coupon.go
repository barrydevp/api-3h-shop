package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOneCoupon(query *connect.QueryMySQL) (*model.Coupon, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, code, discount, description, expires_at, updated_at
		FROM coupons
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

	var _coupon model.Coupon

	err = stmt.QueryRow(args...).Scan(
		&_coupon.RawId,
		&_coupon.RawCode,
		&_coupon.RawDiscount,
		&_coupon.RawDescription,
		&_coupon.RawExpiresAt,
		&_coupon.RawUpdatedAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_coupon.FillResponse()

		return &_coupon, nil
	default:
		return nil, err
	}

	return nil, nil

}
