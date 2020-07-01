package main

import (
	"fmt"
	"unsafe"
)

func voidPointOper() {
	// 类似于C void* 的赋值.
	u := User{1, "Tom"}
	var i interface{} = u // i持有的是目标对象的只读复制品,复制了完整对象.
	var ip interface{} = &u
	u.id = 2
	u.name = "Jack"
	fmt.Printf("u: %v\n", u)
	fmt.Printf("i.(User): %v\n", i.(User)) // 输出仍是 {1 Tom}
	fmt.Printf("ip.(*User): %v\n", ip.(*User))

	// 当使用 interfacr{} 作为void*接收具体类型时，对真实类型的检测
	// 方式1，用switch
	switch i.(type) {
	case *User:
		fmt.Println("It's *User")
	case User:
		fmt.Println("It's User")
	case *Datam:
		fmt.Println("It's *Datam")
	default:
		break
	}
	// 方式2，用if判断
	if vc, ok := i.(User); ok {
		fmt.Println("oh! It's User", vc)
	}

	return
}

func unsafePointOper() {

	// 指针转换
	d := struct {
		s string
		y string
		x int
	}{"abc", "def", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset

	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2)        // Pointer -> *int
	*px = 200               // d.x = 200
	fmt.Printf("%#v\n", d)

	a := [3]int{0, 1, 2}
	as := []int{0, 1, 2}
	fmt.Printf("%T, %T\n", a, as)
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}
	fmt.Println(a)

	switch d.x {
	case 100:
	case 200:
		fmt.Printf("px is %d\n", d.x)
	case 300:
		fmt.Printf("px 300 is %d\n", d.x)

	}

}
