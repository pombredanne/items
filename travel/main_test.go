package main

import "testing"

//func TestGetIPAndPort(t *testing.T) {
//	ip,port :=GetIPAndPort("127.0.0.1:50008")
//	t.Log(ip,port)
//}
func TestCheckEmpty(t *testing.T) {
	type Test struct {
		T1 string
		T2 string
		T3 string
	}
	test :=Test{}
	test.T1="adsf"
	test.T2="2133"
	test.T3="3dfa"
	t.Log(CheckEmpty(test))
}