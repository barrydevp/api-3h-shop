package factories

import (
	"database/sql"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func InsertCategory(insertCategory *model.BodyCategory) (*int64, error) {
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		INSERT categories 
		SET name=?, parent_id=?
	`)

	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	validParentId := sql.NullInt64{Valid: false}

	if insertCategory.ParentId != nil {
		validParentId.Int64 = *insertCategory.ParentId
		validParentId.Valid = true
	}
	res, err := stmt.Exec(&insertCategory.Name, &validParentId)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

