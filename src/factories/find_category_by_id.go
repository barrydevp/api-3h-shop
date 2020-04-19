package factories

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindCategoryById(categoryId int64) (*model.Category, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, name, parent_id, status, updated_at
		FROM categories
		WHERE _id=?
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var _category model.Category

	err = stmt.QueryRow(categoryId).Scan(&_category.RawId, &_category.RawName, &_category.RawParentId, &_category.RawStatus, &_category.RawUpdatedAt)

	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("category does not exists")
	case nil:
		_category.FillResponse()

		return &_category, nil
	default:
		return nil, err
	}

	return nil, nil
}
