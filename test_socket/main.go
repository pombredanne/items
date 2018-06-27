package main

import (
	"sync"
	"test_socket/server"
	"time"
	"test_socket/client"
	"fmt"
)
var waitGroup sync.WaitGroup
func main(){
	waitGroup.Add(2)
	go func(){
		defer waitGroup.Done()
		server.Server()
	}()

	time.Sleep(1*time.Second)
	go func(){
		defer waitGroup.Done()
		client.Client()
	}()
	waitGroup.Wait()
	fmt.Println("finished")
}