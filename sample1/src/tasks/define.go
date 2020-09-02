package tasks

import "runtime"

var (
	GMaxWorker = runtime.NumCPU()
	GMaxQueue  = 64
	GJobQueue  chan Job
)

type Job interface {
	Exec() error
}
