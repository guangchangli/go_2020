package main

import "fmt"

type salaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculatorLeaveLeft() int
}

type employee struct {
	code int
	name string
}

func (e employee) DisplaySalary() {
	fmt.Println("id is %d, name is %s", e.code, e.name)
}
func (e employee) CalculatorLeaveLeft() int {
	return e.code * 2
}
func main() {
	e := employee{
		code: 100,
		name: "liguangchang",
	}
	var l LeaveCalculator = e
	println(l.CalculatorLeaveLeft())
}
