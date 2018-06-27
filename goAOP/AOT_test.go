package main

import (
	"testing"
	"fmt"
)

func TestGetDuration(t *testing.T) {
	t.Log(GetDuration(func() {
		for i:=0;i<1000;i++{
			fmt.Print("do sth")
		}
		fmt.Println("")
	}))
}

func Test(t *testing.T){
	c:=context{
		KV: make(map[string]interface{}),
	}

	wrapper(func(ctx context){
		ctx.KV["return"] = add(1,2)
	},c)

}

func add(a,b int)int{
	return a+b
}

type context struct{
	KV  map[string]interface{}
}
func wrapper (f func(ctx context),c context){
	f(c)
}