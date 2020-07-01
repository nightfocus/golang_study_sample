// ctoola project ctoola.go
package ctoola

import (
	"bytes"
	"runtime"
	"strconv"
)

func Toola(str string) string {

	str1 := str + " <- "
	return str1
}

type MyInfo struct {
	age int16
}

func (mi MyInfo) GetAge() int16 {
	return mi.age
}

func (mi *MyInfo) SetAge(sage int16) {
	mi.age = sage
}

// 获取协程的id
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
