package functionRouter

import "sync"

type  Router struct {
	RMutex sync.Mutex
	Handlers map[string]func()
	RoutineArgs map[string][]chan interface{}
}

func NewRouter() *Router{
	routineArgs := make(map[string][]chan interface{})
	return &Router{ sync.Mutex{},make(map[string]func()),routineArgs}
}

func (r *Router) Add(key string,f func(),){
	r.Handlers[key] = f
}

func (r *Router) SafeAdd(key string,f func()){
	r.RMutex.Lock()
	r.Handlers[key] = f
	r.RMutex.Unlock()
}

func (r *Router) Handle(key string,arg...interface{}){
	r.Handlers[key]()
}