package utils

import (
	"reflect"
	"encoding/json"
)

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

func StructToJsonTagMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Tag.Get("json")
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}

func StructToJsonTagMap2(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &result)

	return result
}



