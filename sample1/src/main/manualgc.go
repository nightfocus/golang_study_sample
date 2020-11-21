package main

/*
  在实际的编程中，我们都希望每个对象释放时执行一个方法，在该方法内执行一些计数、释放或特定的要求，
  以往都是在对象指针置nil前调用一个特定的方法，golang提供了runtime.SetFinalizer函数，
  当GC准备释放对象时，会回调该函数指定的方法，非常方便和有效。

  GC是自动进行的。如果要手动进行GC，可以调用runtime.GC()函数，进行显式GC。
*/

import (
	"fmt"
	//	"log"
	"runtime"
	"sync"
	//	"time"
)

type Car struct {
	Uid   string
	Brand string
}
type CarWrap struct {
	c *Car
}

var (
	mLck     = sync.Mutex{}          // mutex for mCarPort
	mCarPort = make(map[string]*Car) // 用map模型一个car管理容器
)

// 创建一个Car对象给外部（但实际返回的是CarWrap对象），
// 同时在内部用mCarPort记录所有创建的Car对象，在外面用完后自动移除。
func NewCarWrap(inUid string, inBrand string) (*CarWrap, error) {
	nc := &Car{
		Uid:   inUid,
		Brand: inBrand,
	}

	mLck.Lock()
	defer mLck.Unlock()

	// 将新的Car对象，存入到map中
	mCarPort[inUid] = nc
	fmt.Printf("Create car %s, all: %d cars\n", inUid, len(mCarPort))

	cw := &CarWrap{
		c: nc,
	}

	/*
		在这里用cw 来做SetFinalizer，是我们希望在释放时，做些事情。
		例如在这里做的事情就是用FinalizeCarWrap()，将car从map里移除。
		当NewCarWrap() return的cw被使用者用完后（即不再被引用，且调用runtime.GC()后），就触发这个函数。
		但这里不能用nc 来做SetFinalizer，因为nc已存入mCarPort了。所以即使调用者用完，GC也不会释放它。
	*/
	runtime.SetFinalizer(cw, func(cw *CarWrap) {
		FinalizeCarWrap(cw.c.Uid)
	})

	return cw, nil
}

func (*Car) Show() {

}

func FinalizeCarWrap(inUid string) {
	fmt.Printf("Before Finalize car %s, all: %d cars\n", inUid, len(mCarPort))
	mLck.Lock()
	defer mLck.Unlock()
	delete(mCarPort, inUid)
	fmt.Printf("After Finalize car %s, all: %d cars\n", inUid, len(mCarPort))

}

/*
type Road int
type RoadWrap struct {
	rr Road
}

var mmm = make(map[int]*RoadWrap)

func findRoad(r *Road) {

	log.Println("-------road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd

	mmm[0] = &RoadWrap{
		rr: rd,
	}

	runtime.SetFinalizer(r, findRoad)
}

func MMain() {

	entry()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
*/
