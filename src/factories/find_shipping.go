package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindShipping(query *connect.QueryMySQL) ([]*model.Shipping, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, carrier, status, order_id, created_at, updated_at, delivered_at, note, price
		FROM shippings
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
	var listShipping []*model.Shipping

	for rows.Next() {
		_shipping := model.Shipping{}

		err = rows.Scan(
			&_shipping.RawId,
			&_shipping.RawCarrier,
			&_shipping.RawStatus,
			&_shipping.RawOrderId,
			&_shipping.RawCreatedAt,
			&_shipping.RawUpdatedAt,
			&_shipping.RawDeliveredAt,
            &_shipping.RawNote,
            &_shipping.RawPrice,
		)

		if err != nil {
			return nil, err
		}

		_shipping.FillResponse()

		listShipping = append(listShipping, &_shipping)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listShipping, nil
}
