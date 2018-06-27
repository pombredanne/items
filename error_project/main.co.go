package main

import (
	"errors"
)

func  main()  {
	Fun1()
}
type D struct {
	Name string
}
func Fun1() error{

	fun1Error :=errors.New("error made by fun1")
	fun2Error := Fun2()
	return fun1Error
}
func Fun2() error {
	fun2Error :=errors.New("error made by fun2")
	return fun2Error
}
