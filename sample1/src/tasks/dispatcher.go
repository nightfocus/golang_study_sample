package tasks

import (
	. "common"
	"runtime"
	"sync"
)

type Dispatcher struct {
	MaxWorkers int
	WorkerPool chan chan Job
	Quit       chan bool // 控制该Dispatcher退出，但目前没用到
}

func init() {
	GDbgLock = sync.Mutex{}
	runtime.GOMAXPROCS(GMaxWorker)          // 设置逻辑CPU数量
	GJobQueue = make(chan Job, GMaxQueue)   // 初始化带缓冲的chan
	dispatcher := NewDispatcher(GMaxWorker) // 执行NewDispatcher()函数获得一个Dispatcher指针
	dispatcher.Run()
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{MaxWorkers: maxWorkers, WorkerPool: pool, Quit: make(chan bool)}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}

/*
	用法：
type Test struct {
	Name string
}

func (t *Test) Exec() error {
	// time.Sleep(3 * time.Second)
	tasks.DbgPrint("Here is: %s", t.Name)

	return nil
}

type Test2 struct {
	Name string
}

func (t *Test2) Exec() error {
	tasks.DbgPrint("Here is another: %s", t.Name)
	return nil
}

func main(){
	// fetch job
	t := Test{Name: "Coeus"}
	work = &t
	tasks.JobQueue <- work


	t2 := Test2{Name: "Nizz"}
	work = &t2
	tasks.JobQueue <- work
}

*/
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-GJobQueue: // 从JobQueue中接收新的任务
			// DbgPrint("0")
			// jobChannel := <-d.WorkerPool // 从WorkerPool中获得一个已就绪的worker
			// jobChannel <- job            // 将任务分配给这个worker
			// 上面两行，最好是放在一个协程中执行，这样当没有worker就绪时，
			// 或者将这个job写入不了时（jobChannel是不带缓冲的），也不会阻塞这个for循环。
			go func(job Job) {
				DbgPrint("1")
				jobChannel := <-d.WorkerPool // 从WorkerPool中获得一个已就绪的worker
				DbgPrint("2")
				jobChannel <- job // 将任务分配给这个worker
			}(job)

		case <-d.Quit:
			return
		}
	}
}
