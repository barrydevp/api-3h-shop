package factories

import (
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func GetListUser() ([]*model.User, error){
	connection := connections.Mysql.GetConnection()

	stmt, err := connection.Prepare(`
		SELECT
			_id, email, name, password, address, status, created_at, updated_at
		FROM users
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
	var listUser []*model.User

	for rows.Next() {
		_user := model.User{}

		err = rows.Scan(&_user.Id, &_user.Email, &_user.Name, &_user.Password, &_user.Address, &_user.Status, &_user.CreatedAt, &_user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		listUser = append(listUser, &_user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listUser, nil
}
