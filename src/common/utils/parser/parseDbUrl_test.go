package parser

import (
	"testing"
)

func Test_ParseMysqlUrl(t *testing.T) {
	mysqlUrl := "mysql://b08738ff9fff5e:e79a1d81@us-cdbr-iron-east-01.cleardb.net/heroku_e16926abf051efd?reconnect=true"

	if r, e := ParseMysqlUrl(mysqlUrl); e != nil {
		t.Error(e)
	} else {
		t.Log("first test passed")
		t.Log(r)
	}

}
