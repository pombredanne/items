package xormTool

import (
	"github.com/xormplus/xorm"
)
var Db *xorm.Engine
var LocalSession *xorm.Session


func DataSource(dataSource string) (*xorm.Engine,error){
	var err error
	Db,err=xorm.NewPostgreSQL(dataSource)
	if err!=nil{
		return nil,err
	}
	LocalSession = Db.NewSession()
	return Db,nil
}
func Config(printSQL bool,maxIdleConns int,maxOpenConns int){
	//2.显示sql语句
	Db.ShowSQL(printSQL)
	//3.设置连接数
	Db.SetMaxIdleConns(maxIdleConns)
	Db.SetMaxOpenConns(maxOpenConns)

	LocalSession = Db.NewSession()
}
func DefaultConfig(){
	//2.显示sql语句
	Db.ShowSQL(true)

	//3.设置连接数
	Db.SetMaxIdleConns(2000)
	Db.SetMaxOpenConns(1000)


}
func GetDb() *xorm.Engine{
	return Db
}