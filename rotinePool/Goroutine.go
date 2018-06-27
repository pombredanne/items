//To design Wrap() for variable case
package routinePool

import (
	"time"
)

const (
	BLOCK = iota - 1
	FREE  = iota
	BUSY
)
const (
	CONCERNED = iota + 1
	IMPORTANT
	EXTRA
)

type Goroutine struct {
	Id     int64
	Name   string
	Desc   string
	Level  int
	Status int //-1-block,1-free,2-busy
}

//GO(Wrapper(args...)) for those functions with special input args and return args
//GO(f)  for functions without input and return values
func (this *Goroutine) Go(f func()) {
	var done chan byte
	go wrap(f, done)
	<-done
}

//LoopGO(Wrapper(args...)) for those functions with special input args and return args
//LoopGO(f)  for functions without input and return values
func (this *Goroutine) LoopGo(f func(), interval time.Duration) {
	go func() {
		this.Status=BUSY
		for {
			f()
			time.Sleep(interval)
		}
	}()
}
func (this *Goroutine) Destroy() {
	this = nil
}

//to signal done as 1
func wrap(f func(), done chan byte) {
	f()
	done <- 1
}
