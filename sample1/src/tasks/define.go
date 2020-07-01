package tasks

import "runtime"

var (
	GMaxWorker = runtime.NumCPU()
	GMaxQueue  = 512
	GJobQueue  chan Job
)

type Job interface {
	Exec() error
}
