package main

import (
	"fmt"
	"math"
)

// 定义Abser接口类型，里面包含一个Abs函数签名定义，不包括函数实现。
type Abser interface {
	Abs() float64
}

// 定义一个MyFloat类型，它是float64的别名
type MyFloat float64

// 给MyFloat类型，定义一个Abs方法，计算绝对值
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 定义Vertex类型结构体
type Vertex struct {
	X, Y float64
}

// 为Vertex类型定义Abs方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// 定义Abser接口变量
	var abser Abser

	// 定义MyFloat类型变量
	f := MyFloat(-100.25)

	// 因为MyFloat实现了，Abser接口中定义的方法，所以可以赋值给接口类型变量
	abser = f

	// 通过接口类型调用方法
	fmt.Println(abser.Abs())

	// 定义Vertex类型变量
	v := Vertex{10, 20}

	// 将Vertex变量赋值给接口类型，这里之所以需要取地址符，因为*Vertex类型实现了接口的Abs方法
	abser = &v

	// 通过接口类型调用方法
	fmt.Println(abser.Abs())
}
