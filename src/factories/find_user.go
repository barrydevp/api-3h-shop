package factories

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindUser(query *connect.QueryMySQL) ([]*model.User, error) {
	connection := connections.Mysql.GetConnection()

	queryString := `
		SELECT
			_id, email, name, password, address, phone, role, session, status, created_at, updated_at 
		FROM users
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
	var listUser []*model.User

	for rows.Next() {
		_user := model.User{}

		err = rows.Scan(
			&_user.RawId,
			&_user.RawEmail,
			&_user.RawName,
			&_user.RawPassword,
			&_user.RawAddress,
			&_user.RawPhone,
			&_user.RawRole,
			&_user.RawSession,
			&_user.RawStatus,
			&_user.RawCreatedAt,
			&_user.RawUpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		_user.FillResponse()

		listUser = append(listUser, &_user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listUser, nil
}
