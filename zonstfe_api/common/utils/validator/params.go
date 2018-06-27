package validator

import (
	"fmt"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("form")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
		//switch name {
		//case "page":
		//	fields[name].Set(reflect.ValueOf(func(i uint) *uint { return &i }(1)))
		//case "page_size":
		//	fields[name].Set(reflect.ValueOf(func(i uint) *uint { return &i }(20)))
		//}

	}
	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()

				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}

				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}

			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	if v.Kind() != reflect.Ptr {
		return errors.New("not ptr type")
	}

	// 获取指针指向的类型
	vv := v.Type().Elem()
	switch vv.Kind() {
	case reflect.String:
		v.Set(reflect.ValueOf(&value))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		switch vv.Kind() {
		case reflect.Int:
			iv := int(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Int8:
			iv := int8(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Int16:
			iv := int16(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Int32:
			iv := int32(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Int64:
			iv := int64(i)
			v.Set(reflect.ValueOf(&iv))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		switch vv.Kind() {
		case reflect.Uint:
			iv := uint(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Uint8:
			iv := uint8(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Uint16:
			iv := uint16(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Uint32:
			iv := uint32(i)
			v.Set(reflect.ValueOf(&iv))
		case reflect.Uint64:
			iv := uint64(i)
			v.Set(reflect.ValueOf(&iv))
		}
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		switch vv.Kind() {
		case reflect.Float32:
			fv := float32(f)
			v.Set(reflect.ValueOf(&fv))
		case reflect.Float64:
			v.Set(reflect.ValueOf(&f))
		}
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&b))
	case reflect.Map:
		b, err := url.ParseQuery(value)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(&b))
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
