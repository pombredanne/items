package main

import (
	"time"
	"reflect"
	"fmt"
)

//由此方法引申出来的框架,经典模型
//This is the core thought  of AOP framework ,crucial and classical
func GetDuration(f func()) time.Duration{
	t1:=time.Now()
	f()
	t2:=time.Now()
	return t2.Sub(t1)
}

type Context struct{
	ArgMap  map[string]interface{} //存放某一个函数id，对应的入参,interface{}应该合理映射成[]int,[]struct{},[]string...
	OutPutMap map[string]interface{} //存放某一函数id，对应的出参,interface{}应该合理映射成[]int,[]struct{},[]string...
}
type AOPExecutor struct{
}
func (a *AOPExecutor) Register(fun interface{})error{
	vType := reflect.TypeOf(fun)
	er:=check(vType)

}

func check(p reflect.Type)bool{
	_ := p.NumIn()
	defer func() {
		if err := recover(); err != nil {
			
		}
	}()
}
//方法切片
/*如何解析一个方法:
    1.输入
    2.输出
    3.方法名
*/

//索引切片




