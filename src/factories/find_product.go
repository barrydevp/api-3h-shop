package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindProduct(query *connect.QueryMySQL) ([]*model.Product, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, category_id, name, out_price, discount, image_path, description, created_at, updated_at 
		FROM products
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
	var listProduct []*model.Product

	for rows.Next() {
		_product := model.Product{}

		err = rows.Scan(
			&_product.RawId,
			&_product.RawCategoryId,
			&_product.RawOutPrice,
			&_product.RawDiscount,
			&_product.RawImagePath,
			&_product.RawDescription,
			&_product.RawCreatedAt,
			&_product.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_product.FillResponse()

		listProduct = append(listProduct, &_product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listProduct, nil
}
