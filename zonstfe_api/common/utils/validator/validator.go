package validator

import (
	"reflect"
	"strings"
	"fmt"
	"errors"
	"regexp"
	"zonstfe_api/common/options"
	"strconv"
	"encoding/json"
)

var (
	ErrBadParameter = errors.New("bad parameter")
	ErrUnsupported  = errors.New("unsupported type")
)

func Validate(p interface{}) error {
	sv := reflect.ValueOf(p)
	st := reflect.TypeOf(p)
	if sv.Kind() == reflect.Ptr && !sv.IsNil() {
		return Validate(sv.Elem().Interface())
	}
	if sv.Kind() != reflect.Struct && sv.Kind() != reflect.Interface {
		return errors.New("unsupported type")
	}
	size := sv.NumField()
	for i := 0; i < size; i++ {
		f := sv.Field(i)
		tag := st.Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		field_name := strings.ToLower(st.Field(i).Name)
		if st.Field(i).Tag.Get("json") != "" {
			field_name = strings.Split(st.Field(i).Tag.Get("json"), ",")[0]
		}

		tag_list := strings.Split(tag, ",")
		value := reflect.ValueOf(f.Interface())
		if value.Kind() == reflect.Ptr && !value.IsNil() {
			value = reflect.ValueOf(value.Elem().Interface())
		}

		for _, v := range tag_list {
			if strings.HasPrefix(v, "request") {
				if f.IsNil() {
					return errors.New(fmt.Sprintf("%s:must request", field_name))
				}
				// 如果当前值为零值则错误
				if v == "requestz" && value.Interface() == reflect.Zero(f.Type()).Interface() {
					return errors.New(fmt.Sprintf("%s:cant't zero value", field_name))
				}
			} else if strings.HasPrefix(v, "json") {
				if f.IsNil() {
					continue
				}
				if !isJSON(value.String()) {
					return errors.New(fmt.Sprintf("%s:format error", field_name))
				}

			} else if strings.HasPrefix(v, "regexp") {
				if f.IsNil() || value.String() == "" {
					continue
				}
				list := strings.Split(v, "|")
				if len(list) != 2 {
					continue
				}
				if m, _ := regexp.MatchString(regexps[list[1]], value.String()); !m {
					return errors.New(fmt.Sprintf("%s:format error", field_name))
				}
			} else if strings.HasPrefix(v, "option") {
				list := strings.Split(v, "|")
				if len(list) != 2 {
					continue
				}
				if _, ok := options.All[list[1]][value.String()]; !ok {
					return errors.New(fmt.Sprintf("%s:option error", field_name))
				}
			} else if strings.HasPrefix(v, "len") {
				if f.IsNil() {
					continue
				}
				if strings.Contains(v, ">=") {
					list := strings.Split(v, ">=")
					if len(list) != 2 {
						continue
					}
					if err := length(value, ">=", list[1], field_name); err != nil {
						return err
					}
				} else if strings.Contains(v, "<=") {
					list := strings.Split(v, "<=")
					if len(list) != 2 {
						continue
					}
					if err := length(value, "<=", list[1], field_name); err != nil {
						return err
					}
				} else if strings.Contains(v, "!=") {
					list := strings.Split(v, "!=")
					if len(list) != 2 {
						continue
					}
					if err := length(value, "!=", list[1], field_name); err != nil {
						return err
					}
				} else if strings.Contains(v, ">") {
					list := strings.Split(v, ">")
					if len(list) != 2 {
						continue
					}
					if err := length(value, ">", list[1], field_name); err != nil {
						return err
					}

				} else if strings.Contains(v, "<") {
					list := strings.Split(v, "<")
					if len(list) != 2 {
						continue
					}
					if err := length(value, "<", list[1], field_name); err != nil {
						return err
					}
				} else if strings.Contains(v, "==") {
					list := strings.Split(v, "==")
					if len(list) != 2 {
						continue
					}
					if err := length(value, "==", list[1], field_name); err != nil {
						return err
					}
				}

			}
		}

	}
	return nil
}
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}
func length(value reflect.Value, condition, param, field_name string) error {
	invalid := false
	switch value.Kind() {
	case reflect.String:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		switch condition {
		case ">=":
			invalid = int64(len(value.String())) < p
		case "<=":
			invalid = int64(len(value.String())) > p
		case ">":
			invalid = int64(len(value.String())) <= p
		case "<":
			invalid = int64(len(value.String())) >= p
		case "==":
			invalid = int64(len(value.String())) != p
		case "!=":
			invalid = int64(len(value.String())) == p
		}

	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		switch condition {
		case ">=":
			invalid = int64(value.Len()) < p
		case "<=":
			invalid = int64(value.Len()) > p
		case ">":
			invalid = int64(value.Len()) <= p
		case "<":
			invalid = int64(value.Len()) >= p
		case "==":
			invalid = int64(value.Len()) != p
		case "!=":
			invalid = int64(value.Len()) == p

		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := asInt(param)
		if err != nil {
			return ErrBadParameter
		}
		switch condition {
		case ">=":
			invalid = value.Int() < p
		case "<=":
			invalid = value.Int() > p
		case ">":
			invalid = value.Int() <= p
		case "<":
			invalid = value.Int() >= p
		case "==":
			invalid = value.Int() != p
		case "!=":
			invalid = value.Int() == p

		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := asUint(param)
		if err != nil {
			return ErrBadParameter
		}
		switch condition {
		case ">=":
			invalid = value.Uint() < p
		case "<=":
			invalid = value.Uint() > p
		case ">":
			invalid = value.Uint() <= p
		case "<":
			invalid = value.Uint() >= p
		case "==":
			invalid = value.Uint() != p
		case "!=":
			invalid = value.Uint() == p

		}
	case reflect.Float32, reflect.Float64:
		p, err := asFloat(param)
		if err != nil {
			return ErrBadParameter
		}
		switch condition {
		case ">=":
			invalid = value.Float() < p
		case "<=":
			invalid = value.Float() > p
		case ">":
			invalid = value.Float() <= p
		case "<":
			invalid = value.Float() >= p
		case "==":
			invalid = value.Float() != p
		case "!=":
			invalid = value.Float() == p

		}
	default:
		return ErrUnsupported
	}
	if invalid {
		return errors.New(field_name + ":len error")
	}
	return nil
}

// asInt retuns the parameter as a int64
// or panics if it can't convert
func asInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		return 0, ErrBadParameter
	}
	return i, nil
}

// asUint retuns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) (uint64, error) {
	i, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return 0, ErrBadParameter
	}
	return i, nil
}

// asFloat retuns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) (float64, error) {
	i, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0.0, ErrBadParameter
	}
	return i, nil
}
