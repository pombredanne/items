package functionRouter

import (
	"testing"
	"fmt"
	"runtime"
)

func Test(t *testing.T) {
	t.Log("开始测试")
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := NewRouter()
	router.Add("case_1",func(){
		go func(){
			Print()
			}()
	})
	router.Add("case_2",func(){
		Add(1,2)
	})
	router.Handle("case_1")
}


func Print() {
	fmt.Println("Case1: 无返回值，无参数")
}
func Add(a,b int){
	fmt.Println(a+b)
}