package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	// 测试当前计算机对于goroutine是多处理器
	var sw sync.WaitGroup
	sw.Add(10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			defer sw.Done()
			fmt.Println(index)
		}(i)
	}
	sw.Wait()
	fmt.Println("done")
}