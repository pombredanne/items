package main



import (

	"fmt"

	"sync"

	"time"

	"runtime"
)



var m *sync.Mutex



func main() {

	m = new(sync.Mutex)



	go lockPrint(1)

	lockPrint(2)



	time.Sleep(time.Second)



	fmt.Printf("%s\n", "exit!")

}



func lockPrint(i int) {

	println(i, "lock start")

	m.Lock()

	println(i, "in lock")
	time.Sleep()

	m.Unlock()

	println(i, "unlock")
}