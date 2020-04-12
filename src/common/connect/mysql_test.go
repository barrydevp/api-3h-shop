package connect

import (
	"testing"
)

func Test_insertMysql(t *testing.T) {
	mysqlUrl := "mysql://b08738ff9fff5e:e79a1d81@us-cdbr-iron-east-01.cleardb.net/heroku_e16926abf051efd?reconnect=true"

	db := MysqlDB{}

	db.New(mysqlUrl)

	connection := db.GetConnection()

	stmt, err := connection.Prepare("INSERT `users` SET name=?,email=?,password=?,address=?,created_at=NOW()")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := stmt.Exec("barrydevp", "barrydevp@gmail.com", "123456", "Lao Cai")
	if err != nil {
		t.Error(err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("_id inserted", id)
	t.Log("testing passed")
}

func Test_mysql(t *testing.T) {
	mysqlUrl := "mysql://b08738ff9fff5e:e79a1d81@us-cdbr-iron-east-01.cleardb.net/heroku_e16926abf051efd?reconnect=true"

	db := MysqlDB{}

	db.New(mysqlUrl)

	connection := db.GetConnection()

	stmt, err := connection.Prepare("SELECT _id, email, name, password, address, status, created_at, updated_at FROM `users`")
	if err != nil {
		t.Error(err)
		return
	}

	rows, err := stmt.Query()

	if err != nil {
		t.Error(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var (
			_id       int
			email     string
			name      string
			password  string
			address   string
			status    string
			createdAt string
			updatedAt string
		)

		err = rows.Scan(&_id, &email, &name, &password, &address, &status, &createdAt, &updatedAt)

		if err != nil {
			t.Error(err)
			return
		}

		t.Log(_id, email, name, password, address, status, createdAt, updatedAt)
	}

	if err = rows.Err(); err != nil {
		t.Error(err)
		return
	}

	t.Log("testing passed")
}
