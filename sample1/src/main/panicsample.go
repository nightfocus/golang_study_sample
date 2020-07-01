package main

import (
	"fmt"
)

func testPanic() {
	// defer recover() // 无效！
	// defer fmt.Println(recover()) // 无效！
	defer func() {
		func() {
			fmt.Println("defer inner")
			recover() // 无效！
		}()
	}()

	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		// z = x / y
		return
	}()

	/*
		defer func() {
			fmt.Println(recover())
		}()
	*/

	panic("test panic")
}

type Tpa struct {
	CC chan int
}

func testPanic2() {

	var m = Tpa{
		CC: make(chan int),
	}
	close(m.CC)
	//m.CC = nil
	//close(m.CC) // duplicated close will panic

	for {
		v, ok := <-m.CC
		if !ok {
			fmt.Println("m.CC is closed.")
			break
		} else {
			fmt.Println("v: ", v)
		}
	}
}
