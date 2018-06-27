package routinePool

import (
	"time"
	"github.com/pkg/errors"
)

var deFaultPool *GoroutinePool

type GoroutinePool struct {
	MaxRoutine int
	MaxIdle    int
	//gcStrategy string   Unrealized
	RoutineMap map[string]*Goroutine    //key-name value-*routine
	Analyser *Analyser
}

//the goroutinePool is singleton,its optional if you want to make many pools because pool itself is externally Upper-spelled
//Default config through this api as 10000 maxRoutine and 5000 maxIdle ,do Config(xxx,yyy) api if you want change it
func GetRoutinePool() *GoroutinePool {
	if deFaultPool != nil {
		return deFaultPool
	}
	goroutinePool := &GoroutinePool{10000, 5000, make(map[string]*Goroutine, 0),&Analyser{}}
	mainRoutine := &Goroutine{
		Id:    1,
		Name:  "MAIN",
		Desc:  "Basic main routine added in pool by default",
		Level: 0,
	}
	routineGC := &Goroutine{
		Id:    2,
		Name:  "GC",
		Desc:  "a routineGC routine starts by default",
		Level: 0,
	}
	goroutinePool.RoutineMap["MAIN"] = mainRoutine
	goroutinePool.RoutineMap["GC"] = routineGC
	startGC(goroutinePool)
	return goroutinePool
}

//Set maxIdle and max routine number
func (this *GoroutinePool) Config(maxRoutine int, maxIdle int) {
	this.MaxIdle = maxIdle
	this.MaxRoutine = maxRoutine
}

func (this GoroutinePool) GetById(id int64) {
}
func (this GoroutinePool) GetByDesc(desc string) {
}
func (this GoroutinePool) GetByName(desc string) {
}
func (this GoroutinePool) Get(name string) (*Goroutine,error){
	routine := this.RoutineMap[name]
	if routine==nil {
		return this.new(name)
	}
	return this.RoutineMap[name],nil
}
func (this GoroutinePool) new(name string) (*Goroutine,error){
	if len(this.RoutineMap)< this.MaxRoutine {
		goroutine:=&Goroutine{idGenerator(),"","Default",1,FREE}
		return goroutine,nil
	}else{
		return nil,errors.New("MaxRoutine full,please wait for routine dead or set a more maxRoutine number")
	}
}
//clean those nil goroutine
//the real clean work is handled by startGC() whose routine'name is  'GC'
func (this GoroutinePool) Clean() {
	for k,v:=range this.RoutineMap{
		if v==nil{
			delete(this.RoutineMap,k)
		}
	}
}


func idGenerator() int64{
	return time.Now().UnixNano()
}

func startGC(pool *GoroutinePool){
	routine,_:=pool.Get("GC")
	routine.LoopGo(pool.Clean,60*time.Second)
}
