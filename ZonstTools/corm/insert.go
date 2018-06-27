package corm

import (
	"strings"
	"fmt"
	"strconv"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type InsertBuilder struct {
	query  string			//具体的sql语句
	args   []interface{}
	allMap map[string]string   //["sql_columns"]=查询列,["sql_values"]=查询参数,["table_name"]=表名，["returning"]=返回值
}

//确定表名
func Insert(table string) *InsertBuilder {
	return &InsertBuilder{
		allMap: map[string]string{"table_name": table, "returning": ""},
	}
}
//确定列名
func (self *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	self.allMap["sql_columns"] = strings.Join(columns, ",")
	values := make([]string, 0)
	for i, _ := range columns {
		values = append(values, fmt.Sprint("$"+strconv.Itoa(i+1)))
	}
	self.allMap["sql_values"] = strings.Join(values, ",")
	return self
}
//置换allMap形成query语句
func (self *InsertBuilder) Build() {
	self.query = StringMapper(`INSERT INTO {{table_name}} ({{sql_columns}})
	VALUES({{sql_values}})  {{returning}}`, self.allMap)
}
//查询参数填充
func (self *InsertBuilder) Values(value ...interface{}) *InsertBuilder {
	self.args = value
	return self
}

func (self *InsertBuilder) ExecTx(tx *sqlx.Tx) {
	self.Build()
}
func (self *InsertBuilder) Exec() (string, error) {
	self.Build()
	if _, err := dB.Exec(self.query, self.args...); err != nil {
		return "", err
	}
	return ParseSql(self.query, self.args), nil
}

func (self *InsertBuilder) ExecLastId(id string) (string, int64, error) {
	self.allMap["returning"] = "RETURNING " + id
	self.Build()
	var lastInsertId int64
	if err := dB.QueryRow(self.query, self.args...).Scan(&lastInsertId); err != nil {
		return "", 0, err
	}
	return ParseSql(self.query, self.args), lastInsertId, nil
}

func PgCopy(table_name string, args [][]interface{}, columns []string) error {
	tx, err := dB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(pq.CopyIn(table_name, columns...))
	if err != nil {
		return err
	}

	for _, arg := range args {
		_, err = stmt.Exec(arg...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
