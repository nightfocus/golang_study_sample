package tasks

import (
	. "common"
)

type Worker struct {
	workerPool chan chan Job
	jobChannel chan Job
	quit       chan bool
}

func NewWorker(workPool chan chan Job) Worker {
	return Worker{
		workerPool: workPool,
		jobChannel: make(chan Job), // 不带缓冲的chan
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerPool <- w.jobChannel // 告诉WorkerPool，该Worker已经就绪.
			select {
			case job := <-w.jobChannel: // 从chan中接收任务来处理
				DbgPrint("3")
				if err := job.Exec(); err != nil {
					DbgPrint("excute job failed with err: %v", err)
				}
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
