package StructDefaulter

import (
	"reflect"
	"fmt"
)

type Defaulter struct{
}
func (this *Defaulter) Init(input interface{})interface{}{
	vType := reflect.TypeOf(input)
	fmt.Println("1:",vType)
	vValue := reflect.ValueOf(input)
	fmt.Println("2:",vValue.Field(0).String())

	for i := 0; i < vType.NumField(); i++ {
		tagValue :=vType.Field(i).Tag.Get("Defaulter")
		if vValue.Field(i).String()=="undefined"{
			vValue = reflect.ValueOf(&input).Elem()
			vValue.Field(i).SetString(tagValue)
		}
	}
	return input

}

