package routinePool


type Analyser struct{
	RoutineAmount int //amount of routines of a pool
	AliveRoutines []int  //store alive routines' id array
	DeadLockRoutines []int	//store deadlock routines'id array
	HealthStatus int // calculate to a state that stands for a pool's health
}