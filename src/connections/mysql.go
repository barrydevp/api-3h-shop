package connections

import (
	"github.com/barrydev/api-3h-shop/src/common/connect"
	"github.com/barrydev/api-3h-shop/src/constants"
	"log"
)

var Mysql connect.MysqlDB

func init() {
	Mysql = connect.MysqlDB{}

	Mysql.New(constants.CLEARDB_DATABASE_URL)
	log.Println("Running MYSQL ON: ", constants.CLEARDB_DATABASE_URL)
}
