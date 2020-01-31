package main

import (
	"errors"
	"fmt"
	"time"
)

/**
	go error 不会导致程序退出
      panic 会导致程序退出
*/

func main() {
	_, err := do()
	if err != nil {
		fmt.Println(err)
	}

	//测试自定义
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func do() (int, error) {
	return 0, errors.New("error out")
}

type myError struct {
	when time.Time
	what string
}

//实现错误接口方法
func (e *myError) Error() string {
	return fmt.Sprintf("at %v,%s", e.when, e.what)
}

func run() error {
	return &myError{
		when: time.Now(),
		what: "my definition error out",
	}
}
