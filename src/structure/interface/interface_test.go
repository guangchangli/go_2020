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
	f:=myFloat(-100.25)
	inter=f
	fmt.Println(inter.Method())
}
