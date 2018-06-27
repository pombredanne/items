package xormTool

import (
	_ "github.com/lib/pq"
)

/*
	1.插入sql语句和参数，insert返回新增行的id和错误信息
 */
func Insert(sql string,args ...interface{}) (int,error){
	session := Db.NewSession()
	defer session.Close()
	session.Begin()
	var  id int
	_,err:=session.SQL(sql,args...).Get(&id)
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	session.Commit()

	return id,err
}

func LocalSessionInsert(sql string,args ...interface{}) (int,error){
	LocalSession.Begin()
	var  id int
	_,err:=LocalSession.SQL(sql,args...).Get(&id)
	if err!=nil {
		LocalSession.Rollback()
		return -1,nil
	}
	LocalSession.Commit()
	return id,err
}