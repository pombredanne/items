package DAO

import (
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"fmt"
)

var DataSource string
var DB *xorm.Engine

func init() {
	var err error
	//DataSource = "postgres://postgres:123@localhost:5432/travel?sslmode=disable"
	DataSource = "postgres://medium:mediuml4eLxglxL8@111.231.137.127:5432/medium?sslmode=disable"
	DB, err = xorm.NewPostgreSQL(DataSource)
	if err != nil {
		fmt.Println(err)
	}
	//2.显示sql语句
	DB.ShowSQL(true)

	//3.设置连接数
	DB.SetMaxIdleConns(2000)
	DB.SetMaxOpenConns(1000)
}
