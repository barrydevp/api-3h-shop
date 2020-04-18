package connect

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/utils/parser"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlDB struct {
	connection *sql.DB
}

func (db *MysqlDB) New(url string) *MysqlDB {
	validUrl, err := parser.ParseMysqlUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	connection, err := sql.Open("mysql", validUrl)
	if err != nil {
		log.Fatal(err)
	}

	db.connection = connection

	return db
}

func (db *MysqlDB) GetConnection() *sql.DB {
	if db.connection == nil {
		log.Fatal(errors.New("you must call New method before GetConnection"))
	}

	return db.connection
}

func (db *MysqlDB) Close() {
	if db.connection == nil {
		log.Fatal(errors.New("mysqlDB have not init"))
	}

	db.connection.Close()
}

type QueryMySQL struct {
	Where   *string
	Join    *string
	GroupBy *string
	Having  *string
	OrderBy *string
	Limit   *string
	Offset  *string
	Args    []interface{}
}

func (query *QueryMySQL) ToQueryString() string {
	queryString := ""
	if query.Where != nil {
		queryString += "WHERE" + *query.Where
	}

	if query.Join != nil {
		queryString += "\nJOIN" + *query.Join
	}

	if query.GroupBy != nil {
		queryString += "\nGROUP BY" + *query.GroupBy
	}

	if query.Having != nil {
		queryString += "\nHAVING" + *query.Having
	}

	if query.OrderBy != nil {
		queryString += "\nORDER BY" + *query.OrderBy
	}

	if query.Limit != nil {
		queryString += "\nLIMIT" + *query.Limit
	}

	if query.Offset != nil {
		queryString += "\nOFFSET" + *query.Offset
	}

	return queryString
}
