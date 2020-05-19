package factories

import (
	"database/sql"

	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/connections"
	"github.com/barrydev/api-3h-shop/src/model"
)

func FindOneUser(query *connect.QueryMySQL) (*model.User, error) {
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

	var _user model.User

	err = stmt.QueryRow(args...).Scan(
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

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		_user.FillResponse()

		return &_user, nil
	default:
		return nil, err
	}

	return nil, nil
}
