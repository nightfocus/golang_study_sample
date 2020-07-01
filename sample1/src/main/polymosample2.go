package main

import (
	"fmt"
)

type IHelloer interface {
	Hello2(name string)
}

type AA struct {
	s string
}

func (oa AA) Hello2(name string) {
	fmt.Println("In Hello2() : ", name)
}

func (oa AA) Wow() {
	fmt.Println("In Wow()")
}

func Callit(ih IHelloer) {
	ih.Hello2("Yu")
	// 检测ih是否是一个AA类型
	if aa, ok := ih.(AA); ok {
		aa.Wow()
	}
}
