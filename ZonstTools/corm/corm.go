package corm

import (
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"fmt"
	"reflect"
)

var dB *sqlx.DB

func ParseSql(format string, p []interface{}) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "$" + strconv.Itoa(k+1)
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}

func InterfaceSlice(slice interface{}) []interface{} {
	switch slice.(type) {
	case []interface{}:
		return slice.([]interface{})
	default:
		s := reflect.ValueOf(slice)
		if s.Kind() != reflect.Slice {
			panic("InterfaceSlice() given a non-slice type")
		}
		ret := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}
		return ret
	}

}

func GetReflectValue(arg interface{}) interface{} {
	value := reflect.ValueOf(arg)
	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = reflect.ValueOf(value.Elem().Interface())
	}
	return value.Interface()

}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func SetDb(db *sqlx.DB)  {
	dB = db
}
//想自定义raw查询，通过此方法拿到全新的db对象
func GetDb(dataSource string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	return db.Unsafe()
}

func GetCurrentDb() *sqlx.DB {
	return dB
}
func Exec(sql string,args...interface{})error{
	if _, err := dB.Exec(sql,args...); err != nil {
		return  err
	}
	return nil
}
func ExecWithId(sql string,args...interface{}) (int64,error){
	var id int64
	if  err := dB.Get(&id,sql,args...); err != nil {
		return -1, err
	}
	return id,nil
}
func Query(dest interface{},sql string,args ...interface{}) error{
	if  err := dB.Select(dest,sql,args...); err != nil {
		return err
	}
	return nil
}
func QueryOne(dest interface{},sql string,args ...interface{}) error{
	if  err := dB.Get(dest,sql,args...); err != nil {
		return err
	}
	return nil
}
//动态where
func QueryDynamic(dest interface{},basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{}) error{
	//1.处理where
	var sql =basicSql
	if len(whereMap)!=0 {
		sql =sql+" where "
		for i,v:=range whereMap{
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			sql = sql +" "+ v[0]+" "+v[1]+v[2]+"$"+strconv.Itoa(i+1)
		}
	}
	fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby
	if len(orderBy)!=0 && orderBy!=nil{
		sql = sql+" order by "+strings.Join(orderBy,",")+" "+Asc+" "
	}
	fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit!=-1&&offset!=-1{
		sql = sql + " limit "+strconv.Itoa(limit)+" offset "+strconv.Itoa(offset)
	}

	if  err := dB.Select(dest,sql,args...); err != nil {
		return err
	}
	return nil
}


func QueryOneDynamic(dest interface{},basicSql string,whereMap [][]string,orderBy []string,Asc string,limit int,offset int,args...interface{}) error{
	//1.处理where
	var sql =basicSql
	if len(whereMap)!=0 {
		sql =sql+" where "
		for i,v:=range whereMap{
			//v[0]表示性质，and 还是or,v[1]表示field，比如name，age,v[2]表示条件符号,=,>,<,<>,like
			sql = sql +" "+ v[0]+" "+v[1]+v[2]+"$"+strconv.Itoa(i+1)
		}
	}
	fmt.Println("处理where完毕:"+sql)

	//2.处理Orderby
	if len(orderBy)!=0 && orderBy!=nil{
		sql = sql+" order by "+strings.Join(orderBy,",")+" "+Asc+" "
	}
	fmt.Println("处理order,asc完毕:"+sql)

	//3.处理limit,offset
	if limit!=-1&&offset!=-1{
		sql = sql + " limit "+strconv.Itoa(limit)+" offset "+strconv.Itoa(offset)
	}

	if  err := dB.Get(dest,sql,args...); err != nil {
		return err
	}
	return nil
}