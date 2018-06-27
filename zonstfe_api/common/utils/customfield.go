package utils

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/lib/pq/hstore"
	"strings"
	"encoding/json"
)

type JSONFloat float64

func (n JSONFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

type JSONTime struct {
	time.Time
}

func (t *JSONTime) FormatString(str string) string {
	d := time.Time(t.Time).Format(str)
	return d
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	list := strings.Split(fmt.Sprintf("%v", time.Time(t.Time)), " ")
	stamp := fmt.Sprintf("\"%s\"", time.Time(t.Time).Format("2006-01-02 15:04:05"))
	if len(list) >= 2 && list[1] == "00:00:00" {
		stamp = fmt.Sprintf("\"%s\"", time.Time(t.Time).Format("2006-01-02"))
	}
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1: len(b)-1]
	}
	switch len(string(b)) {
	case 10:
		t.Time, err = time.Parse("2006-01-02", string(b))
		return err
	case 16:
		t.Time, err = time.Parse("2006-01-02 15:04", string(b))
		return err
	case 19:
		t.Time, err = time.Parse("2006-01-02 15:04:05", string(b))
		return err
	}
	return nil
}

func (t *JSONTime) Scan(value interface{}) (err error) {
	if value == nil {
		t.Time = time.Now()
		return nil
	}

	var ok bool
	if t.Time, ok = value.(time.Time); !ok {
		return errors.New(fmt.Sprintf("scan %v to JSONTime error.", value))
	}
	return nil
}

// JSONTime for sql.driver.Value.
func (t JSONTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// NullJSONTime represents a time.Time that may be null. NullTime implements the
// sql.Scanner interface so it can be used as a scan destination, similar to
// sql.NullString.
type NullJSONTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (nt *NullJSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(nt.Time).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (nt *NullJSONTime) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		nt.Time, nt.Valid = time.Time{}, false
	}
	if len(b) == 2 && b[0] == '"' && b[1] == '"' {
		nt.Time, nt.Valid = time.Time{}, false
	}

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1: len(b)-1]
	}

	parseTime, err := time.Parse("2006-01-02", string(b))
	nt.Time, nt.Valid = parseTime, true
	return err
}

// Scan implements the Scanner interface.
func (nt *NullJSONTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullJSONTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Postgresql Hstore support
type Hstore struct {
	hstore.Hstore
}

func (h *Hstore) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	for k, v := range h.Hstore.Map {
		m[k] = v.String
	}
	return json.Marshal(m)
}

func (h *Hstore) UnmarshalJSON(b []byte) (err error) {
	m := make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	sqlMap := make(map[string]sql.NullString)
	for k, v := range m {
		sqlMap[k] = sql.NullString{v, true}
	}

	h.Hstore = hstore.Hstore{sqlMap}
	return nil
}

//
// func (h *Hstore) Scan(value interface{}) (err error) {
//     log.Printf("value %T, %v\n", value, value)
//     if value == nil {
//         h.Hstore = hstore.Hstore{}
//         return nil
//     }
//
//     var ok bool
//     if h.Hstore, ok = value.(hstore.Hstore); !ok {
//         return errors.New(fmt.Sprintf("scan %v to JSONTime error.", value))
//     }
//     return nil
// }
//
// func (h *Hstore) Value() (driver.Value, error) {
//     log.Printf("h = %#v\n", h)
//     return driver.Value(&h.Hstore), nil
// }

// ToMap converts a struct to a map using the struct's tags.
//
// ToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			// out[tagv] = v.Field(i).Interface()
			out[tagv] = PtrToElem(v.Field(i).Interface())
		}
	}
	return out, nil
}

func PtrToElem(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return vv.Elem().Interface()
	}
	return v
}

/***************************************************************************/

func Int32tointptr(value int32) *int {
	tmp := int(value)
	return &tmp
}

func Float32toFloat64ptr(value float32) *float64 {
	tmp := float64(value)
	return &tmp
}

func Int64(v int64) *int64 {
	return &v
}

func Int32(v int32) *int32 {
	return &v
}

func Int(v int) *int {
	return &v
}

func String(v string) *string {
	return &v
}

func Float64(v float64) *float64 {
	return &v
}
