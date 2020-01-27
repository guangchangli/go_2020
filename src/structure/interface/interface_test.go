package main

import (
	"fmt"
	"testing"
)

type myFloat float64

type myInter interface {
	Method() float64
}

func (f myFloat) Method() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}

}
func Test_Interface(t *testing.T) {
	//definition interface variable
	var inter myInter
	f := myFloat(-100.25)
	inter = f
	fmt.Println(inter.Method())
}
func Test_NullInterface(t *testing.T) {
	//空接口类似object 可以存储任意类型的数据
	var i interface{}
	i = 5
	fmt.Println(i)
	i = "hello null interface"
	fmt.Println(i)
}
func Test_InterfaceConvert(t *testing.T) {
	//interface 类型转换 t:=i.(T)
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	//判断interface类型转换失败
	s, ok := i.(string)
	fmt.Println(ok)

	//type 只能在 switch 中使用
	switch in := i.(type) {
	case int:
		fmt.Println("int 类型,",in)
	case string:
		fmt.Println("string 类型,",in)
	default:
		fmt.Println("位置类型,",in)
	}
}
