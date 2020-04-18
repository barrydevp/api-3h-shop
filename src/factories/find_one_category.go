package factories

import (
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCategoryById(categoryId *int64) (*model.Category, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, name, parent_id, status, updated_at
		FROM categories
		WHERE _id=?
	`)

	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	var _category model.Category

	err = stmt.QueryRow(categoryId).Scan(&_category.RawId, &_category.RawName, &_category.RawParentId, &_category.RawStatus, &_category.RawUpdatedAt)

	if err != nil {
		return nil, err
	}

	_category.FillResponse()

	return &_category, nil
}
