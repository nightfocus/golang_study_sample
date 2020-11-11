package main

/*
   仿net/http库里的 http.HandleFunc(...)的写法，
   通过HandlerFunc实现了普通函数转一种Handler接口。
*/

import (
	"fmt"
)

type THandler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

// 重点是这里，表示 HandlerFunc 有自己的方法Do，这样就实现了 THandler interface
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[interface{}]interface{}, h THandler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	// 这里传入的f是一个函数，但Each 第二个参数是一个 THandler 类型
	// 所以使用 HandlerFunc(f) 就将 f 由function转换成 THandler 类型了
	Each(m, HandlerFunc(f))
}

func selfInfo(k, v interface{}) {
	fmt.Printf("hello, I am %s, %d age\n", k, v)
}

func TfuncToHandle() {
	persons := make(map[interface{}]interface{})
	persons["zhangsan"] = 20
	persons["lisi"] = 23
	persons["wangwu"] = 26

	EachFunc(persons, selfInfo)

}
