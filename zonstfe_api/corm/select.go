package corm

import (
	"strings"
	"fmt"
	"reflect"
	"strconv"
)

var regex = map[interface{}]string{
	"=":           "",
	">":           "",
	"<":           "",
	"!=":          "",
	"<>":          "",
	">=":          "",
	"<=":          "",
	"like":        "",
	"in":          "",
	"not in":      "",
	"between":     "",
	"not between": "",
}

type SelectBuilder struct {
	sql       string
	argsMap   map[string]interface{}
	allMap    map[string]string
	formDbMap map[string]string
	whereCond []string
	pageSize  uint
	page      uint
}

func Select(sql string) *SelectBuilder {
	return &SelectBuilder{
		sql:       sql,
		argsMap:   make(map[string]interface{}, 0),
		formDbMap: make(map[string]string, 0),
		allMap: map[string]string{
			"sql_where":    "",
			"sql_groupby":  "",
			"sql_order":    "",
			"sql_paging":   "",
			"join_groupby": "",
		},
		page:      1,
		pageSize:  20,
		whereCond: make([]string, 0),
	}
}
func parseEqualWhere(condition string, whereCond *[]string, argsMap map[string]interface{}, formDbMap map[string]string, args ...interface{}) {
	switch len(args) {
	case 3:
		if _, ok := regex[args[1]]; !ok {
			panic("where运算条件参数有误!!")
		}
		if !canEqual(args[2]) {

			argKey := args[0].(string)
			argsMap[argKey] = GetReflectValue(args[2])
			value, ok := formDbMap[argKey]
			if ok {
				*whereCond = append(*whereCond, " "+condition+" "+value+""+args[1].(string)+":"+argKey)
			} else {
				*whereCond = append(*whereCond, " "+condition+" "+argKey+""+args[1].(string)+":"+argKey)
			}
		}
	case 2:
		if !canEqual(args[2]) {
			argKey := args[0].(string)
			argsMap[argKey] = GetReflectValue(args[2])
			value, ok := formDbMap[argKey]
			if ok {
				*whereCond = append(*whereCond, " "+condition+" "+value+"=:"+argKey)

			} else {
				*whereCond = append(*whereCond, " "+condition+" "+argKey+"=:"+argKey)

			}
		}
	case 1:
		switch paramReal := args[0].(type) {
		case string:
			*whereCond = append(*whereCond, condition+" ("+paramReal+")")
		case map[string]interface{}:
			for k, v := range paramReal {
				if !canEqual(v) {
					argsMap[k] = GetReflectValue(v)
					value, ok := formDbMap[k]
					if ok {
						*whereCond = append(*whereCond, " "+condition+" "+value+"=:"+k)
					} else {
						*whereCond = append(*whereCond, " "+condition+" "+k+"=:"+k)

					}
				}
			}
		case [][]interface{}:
			for _, arr := range paramReal {
				parseEqualWhere(condition, whereCond, argsMap, formDbMap, arr...)
			}
		default:
			panic("where条件格式错误")
		}
	}

}
func (self *SelectBuilder) EqualWhere(all string, args ...interface{}) *SelectBuilder {
	condition := "and"
	whereCond := make([]string, 0)
	argsMap := make(map[string]interface{}, 0)
	parseEqualWhere(condition, &whereCond, argsMap, self.formDbMap, args...)
	for k, v := range argsMap {
		self.argsMap[k] = v
	}
	if len(whereCond) > 0 {
		self.allMap[all] = "where " + strings.TrimLeft(strings.Trim(strings.Join(whereCond, " "), " "), "and")
	} else {
		self.allMap[all] = ""
	}
	return self
}

func (self *SelectBuilder) EqualOrWhere(all string, args ...interface{}) *SelectBuilder {
	condition := "or"
	whereCond := make([]string, 0)
	argsMap := make(map[string]interface{}, 0)
	parseEqualWhere(condition, &whereCond, argsMap, self.formDbMap, args...)
	for k, v := range argsMap {
		self.argsMap[k] = v
	}
	if len(whereCond) > 0 {
		self.allMap[all] = "where " + strings.TrimLeft(strings.Trim(strings.Join(whereCond, " "), " "), "and")
	} else {
		self.allMap[all] = ""
	}
	return self
}

func canEqual(ptr interface{}) bool {
	v := reflect.ValueOf(ptr)
	return v.IsNil()
}

func (self *SelectBuilder) SetAllMap(allMap map[string]string) *SelectBuilder {
	for k, v := range allMap {
		self.allMap[k] = v
	}
	return self
}

func (self *SelectBuilder) Loads(dest interface{}) error {
	self.parseAll()
	sql := StringMapper(fmt.Sprintf(`%s {{sql_order}} {{sql_paging}}`, self.sql), self.allMap)
	query, args := self.whereMapper(sql, self.argsMap)
	err := Db.Select(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (self *SelectBuilder) Load(dest interface{}) error {
	self.parseAll()
	sql := StringMapper(fmt.Sprintf(`%s {{sql_order}} {{sql_paging}}`, self.sql), self.allMap)
	query, args := self.whereMapper(sql, self.argsMap)
	fmt.Println(query, args)
	err := Db.Get(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (self *SelectBuilder) Paginate(page, page_size *uint) *SelectBuilder {
	value := fmt.Sprintf(" LIMIT %d OFFSET %d ", self.pageSize, self.pageSize*(self.page-1))
	if page != nil && page_size != nil {
		value = fmt.Sprintf(" LIMIT %d OFFSET %d ", *page_size, *page_size*(*page-1))
	} else if page != nil && page_size == nil {
		value = fmt.Sprintf(" LIMIT %d OFFSET %d ", self.pageSize, self.pageSize*(*page-1))
	}
	self.allMap["sql_paging"] = value
	return self
}

func StringMapper(format string, p map[string]string) string {
	args, i := make([]string, len(p)*2), 0
	for k, v := range p {
		args[i] = "{{" + k + "}}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}
func (self *SelectBuilder) whereMapper(format string, p map[string]interface{}) (string, []interface{}) {
	args, i, ii := make([]string, len(p)*2), 0, 1
	args2 := make([]interface{}, 0)
	for k, v := range p {
		xx := reflect.ValueOf(v)
		t := Deref(xx.Type())
		args[i] = ":" + k
		if t.Kind() == reflect.Slice {
			xx_list := make([]string, 0)
			for index := 0; index < xx.Len(); index++ {
				xx_list = append(xx_list, fmt.Sprint("$"+strconv.Itoa(ii)))
				ii += 1
			}
			args[i+1] = strings.Join(xx_list, ",")
			args2 = append(args2, InterfaceSlice(v)...)

		} else {
			args[i+1] = fmt.Sprint("$" + strconv.Itoa(ii))
			ii += 1
			args2 = append(args2, v)
		}
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format), args2
}

func (self *SelectBuilder) Total(total *int64) error {
	count_string := StringMapper(fmt.Sprintf(`select count(*) from (%s) as t1000`, self.sql), self.allMap)
	querycount, args := self.whereMapper(count_string, self.argsMap)
	err := Db.Get(total, querycount, args...)
	return err
}

func (self *SelectBuilder) parseWhere(condition string, args ...interface{}) *SelectBuilder {
	switch len(args) {
	case 3: // 常规3个参数:  {"id",">",1}
		if _, ok := regex[args[1]]; !ok {
			panic("where运算条件参数有误!!")
		}
		argKey := args[0].(string)
		switch args[1].(string) {
		case "like":
			self.argsMap[argKey] = args[2]
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" like %:"+argKey+"%")
		case "in":
			self.argsMap[argKey] = InterfaceSlice(args[2])
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" in (:"+argKey+")")
		case "not in":
			self.argsMap[argKey] = InterfaceSlice(args[2])
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+" not in (:"+argKey+")")
		default:
			self.argsMap[argKey] = args[2]
			self.whereCond = append(self.whereCond, " "+condition+" "+argKey+""+args[1].(string)+":"+argKey)
		}
	case 2: // 常规2个参数:  {"id",1}
		argKey := args[0].(string)
		self.argsMap[argKey] = args[2]
		self.whereCond = append(self.whereCond, " "+condition+" "+argKey+"=:"+argKey)
	case 1: // 二维数组或字符串
		switch paramReal := args[0].(type) {
		case string:
			self.whereCond = append(self.whereCond, condition+" ("+paramReal+")")
		case map[string]interface{}:
			for k, v := range paramReal {
				self.argsMap[k] = v
				self.whereCond = append(self.whereCond, " "+condition+" "+k+"=:"+k)
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

func (self *SelectBuilder) Where(args ...interface{}) *SelectBuilder {
	return self.parseWhere("and", args...)
}
func (self *SelectBuilder) OrWhere(args ...interface{}) *SelectBuilder {
	return self.parseWhere("or", args...)
}

func (self *SelectBuilder) Req(req interface{}) *SelectBuilder {
	sv := reflect.ValueOf(req)
	st := reflect.TypeOf(req)
	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return self.Req(sv.Elem().Interface())
	}
	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		panic("unsupported type")
	}
	size := sv.NumField()
	for i := 0; i < size; i++ {
		f := sv.Field(i)
		// 判断是否为空
		if f.IsNil() {
			continue
		}
		queryTag := st.Field(i).Tag.Get("query")
		formTag := st.Field(i).Tag.Get("form")
		dbTag := st.Field(i).Tag.Get("db")
		if queryTag == "" || formTag == "" || dbTag == "" {
			continue
		}
		self.formDbMap[formTag] = dbTag
		value := reflect.ValueOf(f.Interface())
		if value.Kind() == reflect.Ptr && !value.IsNil() {
			value = reflect.ValueOf(value.Elem().Interface())
		}
		self.argsMap[formTag] = value.Interface()
		switch queryTag {
		case "groupby":
			self.allMap["groupby"] = value.String()
		case "order":
			if self.allMap["sql_order"] != "" {
				self.allMap["sql_order"] = self.allMap["sql_order"] + ", " + value.String()
			} else {
				self.allMap["sql_order"] = value.String()
			}
		case "orderdesc":
			sep := strings.Split(value.String(), ",")
			for _, s := range sep {
				if self.allMap["sql_order"] == "" {
					self.allMap["sql_order"] = self.allMap["sql_order"] + s + " desc"
				} else {
					self.allMap["sql_order"] = self.allMap["sql_order"] + ", " + s + " desc"
				}
			}
		case "eq":
			self.whereCond = append(self.whereCond, " and "+dbTag+"=:"+formTag)
		case "neq":
			self.whereCond = append(self.whereCond, " and "+dbTag+"!=:"+formTag)
		case "gt":
			self.whereCond = append(self.whereCond, " and "+dbTag+">:"+formTag)
		case "gte":
			self.whereCond = append(self.whereCond, " and "+dbTag+">=:"+formTag)
		case "lt":
			self.whereCond = append(self.whereCond, " and "+dbTag+"<:"+formTag)
		case "lte":
			self.whereCond = append(self.whereCond, " and "+dbTag+"<=:"+formTag)
		case "like":
			self.whereCond = append(self.whereCond, " and "+dbTag+"like %:"+formTag+"%")
		case "in":
			in_list := strings.Split(value.String(), ",")
			if len(in_list) < 0 {
				panic("empty slice passed to 'in' query")
			}
			self.whereCond = append(self.whereCond, " and "+dbTag+" in (:"+formTag+")")
			self.argsMap[formTag] = InterfaceSlice(in_list)

		case "nin":
			nin_list := strings.Split(value.String(), ",")
			if len(nin_list) < 0 {
				panic("empty slice passed to 'nin' query")
			}
			self.whereCond = append(self.whereCond, " and "+dbTag+" not in (:"+formTag+")")
			self.argsMap[formTag] = InterfaceSlice(nin_list)
		default:
			continue
		}
	}
	return self
}

func (self *SelectBuilder) parseAll() *SelectBuilder {
	if self.allMap["sql_order"] != "" {
		self.allMap["sql_order"] = "order by " + self.allMap["sql_order"]
	}
	if self.allMap["sql_groupby"] != "" {
		self.allMap["join_groupby"] = "," + self.allMap["sql_groupby"]
		self.allMap["sql_groupby"] = "group by " + self.allMap["sql_groupby"]
	}
	if len(self.whereCond) > 0 {
		self.allMap["sql_where"] = "where " + strings.TrimLeft(strings.Trim(strings.Join(self.whereCond, " "), " "), "and")
	} else {
		self.allMap["sql_where"] = ""

	}
	if _, ok := self.allMap["sql_paging"]; !ok {
		self.allMap["sql_paging"] = ""
	}
	return self
}
