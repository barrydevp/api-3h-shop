package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCategory(query *connect.QueryMySQL) ([]*model.Category, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, name, image_path, parent_id, status, updated_at
		FROM categories
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
	var listCategory []*model.Category

	for rows.Next() {
		_category := model.Category{}

		err = rows.Scan(
			&_category.RawId,
			&_category.RawName,
			&_category.ImagePath,
			&_category.RawParentId,
			&_category.RawStatus,
			&_category.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_category.FillResponse()

		listCategory = append(listCategory, &_category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listCategory, nil
}
