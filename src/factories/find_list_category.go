package factories

import (
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindListCategory() ([]*model.Category, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, name, parent_id, status, updated_at
		FROM categories
	`)

	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listCategory []*model.Category

	for rows.Next() {
		_category := model.Category{}

		err = rows.Scan(&_category.RawId, &_category.RawName, &_category.RawParentId, &_category.RawStatus, &_category.RawUpdatedAt)

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
