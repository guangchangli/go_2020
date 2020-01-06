package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Random() int {
	return int(C.random())

}
//func Seed(i int) int {
//	return int(C.srandom(C.uint(i)))
//}
func main() {
	var i = 1
	println(C.uint(i))
	fmt.Println(Random())
	//fmt.Println(Seed(5))
	Print("123")
}

func Print(s string) {
	cs := C.CString(s)
	//手动释放内存
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
