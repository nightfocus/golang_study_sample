package tasks

import (
	. "common"
)

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	Quit       chan bool
}

func NewWorker(workPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workPool,
		JobChannel: make(chan Job), // 不带缓冲的chan
		Quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel // 告诉WorkerPool，该Worker已经就绪.
			select {
			case job := <-w.JobChannel: // 从chan中接收任务来处理
				DbgPrint("3")
				if err := job.Exec(); err != nil {
					DbgPrint("excute job failed with err: %v", err)
				}
			case <-w.Quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
