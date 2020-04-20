package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindProductById(productId int64) (*model.Product, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, category_id, name, out_price, discount, image_path, description, created_at, updated_at 
		FROM products
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _product model.Product

	err = stmt.QueryRow(productId).Scan(
		&_product.RawId,
		&_product.RawCategoryId,
		&_product.RawName,
		&_product.RawOutPrice,
		&_product.RawDiscount,
		&_product.RawImagePath,
		&_product.RawDescription,
		&_product.RawCreatedAt,
		&_product.RawUpdatedAt,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_product.FillResponse()

		return &_product, nil
	default:
		return nil, err
	}

	return nil, nil
}
