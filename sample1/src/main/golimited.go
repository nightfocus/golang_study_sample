package main

/*
	按允许最大并发数创建一个带缓冲的通道。
	创建协程之前调用Add()往通道里写一个数据, 	协程完成是调用Done()方法读取一个数据.
	若无法往通道里写数据时, 表示通道已经写满, 也就是目前的协程并发数为允许的最大数量.
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
