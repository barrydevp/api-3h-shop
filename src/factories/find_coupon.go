package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCoupon(query *connect.QueryMySQL) ([]*model.Coupon, error) {
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

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listCoupon []*model.Coupon

	for rows.Next() {
		_coupon := model.Coupon{}

		err = rows.Scan(
			&_coupon.RawId,
			&_coupon.RawCode,
			&_coupon.RawDiscount,
			&_coupon.RawDescription,
			&_coupon.RawExpiresAt,
			&_coupon.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_coupon.FillResponse()

		listCoupon = append(listCoupon, &_coupon)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listCoupon, nil
}
