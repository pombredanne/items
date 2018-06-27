package corm

import (
	"fmt"
	"strconv"
	"strings"
)

type UpdateBuilder struct {
	query     string
	args      []interface{}
	position  int
	allMap    map[string]string
	whereCond []string
}

func Update(table string) *UpdateBuilder {
	return &UpdateBuilder{
		position: 1,
		allMap:   map[string]string{"table_name": table, "returning": ""},
	}
}
func (self *UpdateBuilder) Build() {
	if len(self.whereCond) > 0 {
		self.allMap["sql_where"] = strings.TrimLeft(strings.Trim(strings.Join(self.whereCond, " "), " "), "and")
	} else {
		self.allMap["sql_where"] = ""
	}
	// 如果where 为空会报错 防止批量更新
	self.query = StringMapper(`UPDATE {{table_name}}  SET {{sql_set}} WHERE {{sql_where}}  {{returning}}`, self.allMap)
}

func (self *UpdateBuilder) Set(set map[string]interface{}) *UpdateBuilder {
	sql_set := make([]string, 0)
	for col, v := range set {
		sql_set = append(sql_set, col+" = "+fmt.Sprint("$"+strconv.Itoa(self.position)))
		self.position += 1
		self.args = append(self.args, v)
	}
	self.allMap["sql_set"] = strings.Join(sql_set, ", ")
	return self
}

func (self *UpdateBuilder) Exec() (string, error) {
	self.Build()
	fmt.Println(self.query, self.args)
	if _, err := Db.Exec(self.query, self.args...); err != nil {
		return "", err
	}
	return ParseSql(self.query, self.args), nil
}

func (self *UpdateBuilder) ExecLastId(id string) (string, int64, error) {
	self.allMap["returning"] = "RETURNING " + id
	self.Build()
	var lastInsertId int64
	if err := Db.QueryRow(self.query, self.args...).Scan(&lastInsertId); err != nil {
		return "", 0, err
	}
	return ParseSql(self.query, self.args), lastInsertId, nil
}

func (self *UpdateBuilder) Where(args ...interface{}) *UpdateBuilder {
	return self.parseWhere("and", args...)
}
func (self *UpdateBuilder) OrWhere(args ...interface{}) *UpdateBuilder {
	return self.parseWhere("or", args...)
}

func (self *UpdateBuilder) parseWhere(condition string, args ...interface{}) *UpdateBuilder {
	switch len(args) {
	case 3: // 常规3个参数:  {"id",">",1}
		if _, ok := regex[args[1]]; !ok {
			panic("where运算条件参数有误!!")
		}
		argKey := args[0].(string)
		switch args[1].(string) {
		case "like":
			self.args = append(self.args, args[2])
			self.position += 1
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" like %"+fmt.Sprint("$"+strconv.Itoa(self.position))+"%")
		case "in":
			in_list := InterfaceSlice(args[2])
			if len(in_list) <= 0 {
				panic("empty slice passed to 'in' query")
			}
			self.args = append(self.args, in_list...)
			in_arg := make([]string, 0)
			for i := 0; i < len(in_list); i++ {
				in_arg = append(in_arg, fmt.Sprint("$"+strconv.Itoa(self.position)))
				self.position += 1
			}
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" in ("+strings.Join(in_arg, ",")+")")
		case "not in":
			in_list := InterfaceSlice(args[2])
			if len(in_list) <= 0 {
				panic("empty slice passed to 'in' query")
			}
			self.args = append(self.args, in_list...)
			in_arg := make([]string, 0)
			for i := 0; i < len(in_list); i++ {
				in_arg = append(in_arg, fmt.Sprint("$"+strconv.Itoa(self.position)))
				self.position += 1
			}
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" not in ("+strings.Join(in_arg, ",")+")")
		default:
			self.args = append(self.args, args[2])
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+""+args[1].(string)+fmt.Sprint("$"+strconv.Itoa(self.position)))
			self.position += 1
		}
	case 2: // 常规2个参数:  {"id",1}
		argKey := args[0].(string)
		self.args = append(self.args, args[2])
		self.whereCond = append(self.whereCond, " "+condition+" "+argKey+"="+fmt.Sprint("$"+strconv.Itoa(self.position)))
		self.position += 1
	case 1: // 二维数组或字符串
		switch paramReal := args[0].(type) {
		case string:
			self.whereCond = append(self.whereCond, condition+" ("+paramReal+")")
		case map[string]interface{}:
			for k, v := range paramReal {
				self.args = append(self.args, v)
				self.whereCond = append(self.whereCond, " "+condition+" "+k+"="+fmt.Sprint("$"+strconv.Itoa(self.position)))
				self.position += 1
			}
		case [][]interface{}:
			for _, arr := range paramReal { // {{"a", 1}, {"id", ">", 1}}
				self.parseWhere(condition, arr...)
			}
		default:
			panic("where条件格式错误")
		}
	}
	return self
}
