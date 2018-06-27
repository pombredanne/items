package corm

import (
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"fmt"
	"reflect"
)

var Db *sqlx.DB

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
