package connections

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"log"
	"os"
)

var Mysql connect.MysqlDB

func init() {
	const DefaultMysqlUrl = "mysql://b08738ff9fff5e:e79a1d81@us-cdbr-iron-east-01.cleardb.net/heroku_e16926abf051efd?reconnect=true"

	mysqlUrl := os.Getenv("CLEARDB_DATABASE_URL")
	if mysqlUrl == "" {
		log.Println("$CLEARDB_DATABASE_URL must be set => Using default CLEARDB_DATABASE_URL: ", DefaultMysqlUrl)
		mysqlUrl = DefaultMysqlUrl
	}

	Mysql = connect.MysqlDB{}

	Mysql.New(mysqlUrl)

	log.Println("MYSQL_URL: ", mysqlUrl)
}


