package main

import "fmt"

/**
panic 类似throw 没有拦截就退出
*/
func main() {
	fmt.Println("a")
	panic("error aa")
	fmt.Println("b")
}

