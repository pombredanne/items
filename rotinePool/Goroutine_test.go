package routinePool

import (
	"testing"
	"runtime"
	"fmt"
	"time"
)

func TestStatusLevel(t *testing.T){
	t.Log(BLOCK,FREE,BUSY)
	t.Log(CONCERNED,IMPORTANT,EXTRA)
}

func TestGoroutine_Go(t *testing.T) {
	goroutine := Goroutine{}
	goroutine.Go(Wrapper(3,1.7))
	goroutine.Go(Print)
	runtime.Gosched()
}

func TestGoroutine_LoopGo(t *testing.T) {
	goroutine := Goroutine{}
	goroutine.LoopGo(Wrapper(3,1.7),0)
	goroutine.LoopGo(Print,0)
	time.Sleep(10*time.Second)
}
func Add(a int ,b float32) float32{
	fmt.Println(float32(a)+b)
	return float32(a)+b
}
func Print(){
	fmt.Println("RotinePool")
}
func Wrapper(a int ,b float32) func(){
	return func(){
		Add(a,b)
	}
}

func BenchmarkGoroutine_Go(b *testing.B) {
	for i:=0;i<b.N;i++{
		goroutine := Goroutine{}
		goroutine.Go(Wrapper(3,1.7))
		goroutine.Go(Print)
		runtime.Gosched()
	}
}
