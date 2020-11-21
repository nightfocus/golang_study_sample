package main

import (
	"sample1/src/common"

	"time"
)

// 最最简单的定时器
func ezTimer() {
	// 实现每隔3秒执行一个任务
	//*
	c3 := make(chan int)
	for {
		select {
		case <-time.After(3 * time.Second):
			common.DbgPrint("This is a mission per 3 seconds.\n")
		case <-time.After(800 * time.Millisecond): // 0.8秒触发一次，因为它的存在，所以前面的3秒永远不会被触发
			common.DbgPrint("This is a mission per 0.8 seconds.\n")
		case <-c3: //  如果这里被触发，那么前面的time.After()会重新计时
			common.DbgPrint("recived c3\n")
		}
	}
	//*/
}
