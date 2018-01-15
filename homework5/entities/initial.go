package entities

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mydb *xorm.Engine

// InitMydb .
func InitMydb(name string, password string, port string, dname string) {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	sql := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8&parseTime=true", name, password, port, dname)
	db, err := xorm.NewEngine("mysql", sql)
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:2048)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	// 同步注册表
	err = db.Sync(new(UserInfo))
	if err != nil {
		panic(err)
	}

	mydb = db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
