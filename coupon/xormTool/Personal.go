package xormTool
//订制特殊查询
import (
	"github.com/xormplus/xorm"
)

var hasError = true
//特殊需求:新启用session查询，并且用完即关
func Exec(sql string ,args...interface{}){
	session:=Db.NewSession()
	defer session.Close()
	session.Begin()

	//特殊需求
	if hasError {
		session.Rollback()
	}
	session.Commit()
}

//使用全局session查询
func ExecLocalSession(sql string ,args...interface{}){
	session:=LocalSession
	session.Begin()

	//特殊需求

	if hasError {
		session.Rollback()
	}
	session.Commit()
}

//新起用一个session，并且长线使用
func ExecPeriodSession(sql string ,args...interface{}) *xorm.Session{
	session:=Db.NewSession()
	session.Begin()

	//特殊需求

	if hasError {
		session.Rollback()
	}
	session.Commit()
	return session
}