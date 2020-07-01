// +build linux

package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static char* cat(char* str1, char* str2) {
	static char buf[256];
	strcpy(buf, str1);
	strcat(buf, str2);

	return buf;
}
*/
import "C"

func cgosample() {
	// 使用cgo直接操作C语言的struct
	// 只能在Linux下使用.
	//*
	str1, str2 := "hello3", " world4"
	// golang string -> c string
	cstr1, cstr2 := C.CString(str1), C.CString(str2)
	defer C.free(unsafe.Pointer(cstr1)) // must call
	defer C.free(unsafe.Pointer(cstr2))
	cstr3 := C.cat(cstr1, cstr2)
	// c string -> golang string
	str3 := C.GoString(cstr3)
	fmt.Println(str3) // "hello3 world4"
	//*/
}
