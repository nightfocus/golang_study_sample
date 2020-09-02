package main

/*
	按允许最大并发数创建一个带缓冲的通道。
	创建协程之前调用Add()往通道里写一个数据, 	协程函数进入后，立即调用Done()方法。

	若无法往通道里写数据时, 也就是通道满了，达到目前的协程并发最大数，
	Add()方法将被阻塞, 也就无法创建新的协程.	直到有协程运行完成, 调用Done()方法读取了通道了一个数据.
*/
type GoLimit struct {
	ch chan struct{}
}

func NewGoLimit(max int) *GoLimit {
	return &GoLimit{ch: make(chan struct{}, max)}
}

func (g *GoLimit) Add() {
	g.ch <- struct{}{}
}

func (g *GoLimit) Done() {
	<-g.ch
}
