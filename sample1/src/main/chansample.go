package main

import (
	"fmt"
	"sync"
	"time"
)

func chanSample() {
	aa, bb := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			// 随机选择可用 channel，接收数据。
			select {
			case v, ok = <-aa:
				if ok {
					s = "a1"
					fmt.Println(s, v)
					time.Sleep(1000 * time.Millisecond)
				} else {
					aa = nil // aa is closed, 置为nil,否则会死循环，一直进入到这里.
				}
			case v, ok = <-bb:
				if ok {
					s = "b"
					fmt.Println(s, v)
				} else {
					fmt.Println("ooh! b is closed.")
					bb = nil
				}

			}
		}
	}()

	go func() {
		for aa != nil {
			v, ok := <-aa
			if ok {
				s := "a2"
				fmt.Println(s, v)
				time.Sleep(1000 * time.Millisecond)
			} else {
				aa = nil // aa is closed, 置为nil,否则会死循环，一直进入到这里.
			}
		}
	}()

	aa <- 0
	aa <- 1
	aa <- 2
	bb <- 10
	bb <- 11
	close(bb) // 关闭bb后触发什么呢
	aa <- 3
	aa <- 4
	aa <- 5
}

// channel 同步通知退出机制
func chanSyncSample() {

	var wg sync.WaitGroup
	quit := make(chan bool)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := func() {
				fmt.Println("In chanSyncSample() task(): ", id, time.Now().Nanosecond())
				time.Sleep(time.Second)
				fmt.Println("In chanSyncSample() task() END.")
			}
			for {
				select {
				case <-quit: // closed channel 不会阻塞，因此可用作退出通知。
					fmt.Println("chanSyncSample() END.")
					return
				default: // 执行正常任务。
					task()
				}
			}
		}(i)
	}

	time.Sleep(time.Second * 2) // 让测试 goroutine 运行一会。
	close(quit)                 // 发出退出通知。
	wg.Wait()                   // 等待wg调用了对等Add()次数的Done()操作后，返回.

	return
}
