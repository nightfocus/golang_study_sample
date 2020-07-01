package main

import (
	"fmt"
)

type IHello interface {
	Hello2(name string)
}

type A struct {
	IHello
}

func (*A) Hello2(name string) {
	fmt.Println("hello2 " + name + ", i am IHello.Hello2")
}

func (*A) Hello(name string) {
	fmt.Println("hello " + name + ", i am a")
}

type B struct {
	ima *A
}

func (*B) Hello(name string) {
	fmt.Println("hello " + name + ", i am b")
}

type Hello struct {
}

func (h *Hello) Myprint(str string) {
	fmt.Println("this is point a callback,", str)
	return
}

/*
func (h Hello) Myprint(str string) {
	fmt.Println("this is instance a callback,", str)
	return
}
*/

type MyCallbacker interface {
	Myprint(str string)
	// Froth() int
}

func testCallback(c MyCallbacker) {
	c.Myprint("hello")

	defer func() {
		fmt.Println("In testCallback defer func()")
	}()

	fmt.Println("this is a test callback")
}
