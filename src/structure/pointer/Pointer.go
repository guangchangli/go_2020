package main

import (
	"runtime"
	"time"
)

func main() {
	i := 123
	var p *int
	p = &i
	println(*p)
	println(time.Now().Weekday())
	println(runtime.GOOS)
	println(time.Now().Hour())
	//println(time.Now().Date())
	//println(time.Now().Clock())
	println(time.Now().Day())

	defer println("hello")
	println("world")
	testDefer()
}

func testDefer(){
	for i:=0;i<3;i++{
		defer println(i)
	}
	println("start loop")
}
