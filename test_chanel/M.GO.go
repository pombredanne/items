package main

import (
	"fmt"
	"runtime"
)

func main() {
	flag := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("完成任务ing %d", i)
			for true {
				select {
				case <-flag:
					break
				default:
					continue
				}
			}
			fmt.Printf("任务%d结束", i)
		}()
	}
	flag <- true
	runtime.Gosched()
}
