package corm

import (
	"strings"
	"fmt"
	"strconv"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type InsertBuilder struct {
	query  string
	args   []interface{}
	allMap map[string]string
}

func Insert(table string) *InsertBuilder {
	return &InsertBuilder{
		allMap: map[string]string{"table_name": table, "returning": ""},
	}
}
func (self *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	self.allMap["sql_columns"] = strings.Join(columns, ",")
	values := make([]string, 0)
	for i, _ := range columns {
		values = append(values, fmt.Sprint("$"+strconv.Itoa(i+1)))
	}
	self.allMap["sql_values"] = strings.Join(values, ",")
	return self
}
func (self *InsertBuilder) Build() {
	self.query = StringMapper(`INSERT INTO {{table_name}} ({{sql_columns}})
	VALUES({{sql_values}})  {{returning}}`, self.allMap)
}
func (self *InsertBuilder) Values(value ...interface{}) *InsertBuilder {
	self.args = value
	return self
}

func (self *InsertBuilder) ExecTx(tx *sqlx.Tx) {
	self.Build()
}
func (self *InsertBuilder) Exec() (string, error) {
	self.Build()
	if _, err := Db.Exec(self.query, self.args...); err != nil {
		return "", err
	}
	return ParseSql(self.query, self.args), nil
}

func (self *InsertBuilder) ExecLastId(id string) (string, int64, error) {
	self.allMap["returning"] = "RETURNING " + id
	self.Build()
	var lastInsertId int64
	if err := Db.QueryRow(self.query, self.args...).Scan(&lastInsertId); err != nil {
		return "", 0, err
	}
	return ParseSql(self.query, self.args), lastInsertId, nil
}

func PgCopy(table_name string, args [][]interface{}, columns []string) error {
	tx, err := Db.Begin()
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
