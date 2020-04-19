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
	if query.Where != nil && *query.Where != "" {
		queryString += "WHERE" + *query.Where + "\n"
	}

	if query.Join != nil && *query.Join != "" {
		queryString += "JOIN" + *query.Join + "\n"
	}

	if query.GroupBy != nil && *query.GroupBy != "" {
		queryString += "GROUP BY" + *query.GroupBy + "\n"
	}

	if query.Having != nil && *query.Having != "" {
		queryString += "HAVING" + *query.Having + "\n"
	}

	if query.OrderBy != nil && *query.OrderBy != "" {
		queryString += "ORDER BY" + *query.OrderBy + "\n"
	}

	if query.Limit != nil && *query.Limit != "" {
		queryString += "LIMIT"
		if query.Offset != nil && *query.Offset != "" {
			queryString += *query.Offset + ","
		}
		queryString += *query.Limit + "\n"
	}


	return queryString
}
