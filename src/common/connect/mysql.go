package connect

import (
	"database/sql"
	"errors"
	"github.com/barrydev/api-3h-shop/src/common/utils/parser"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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

	if err = connection.Ping(); err != nil {
		log.Fatal(err)
	}

	connection.SetConnMaxLifetime(3595 * time.Second)
	connection.SetMaxIdleConns(0)

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
	QueryString string
	Args        []interface{}
}
