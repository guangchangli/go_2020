package main

import "fmt"

func main() {
	//initMap()
	addElement()
}

func initMap() {

	//key int value string
	var m map[int]string
	fmt.Println(m)
	n := make(map[int]string)
	fmt.Println(n)
	mm := make(map[string]string)
	fmt.Println(mm)
}

func addElement() {
	m := make(map[string]string)
	m["name"] = "liguangchang"
	fmt.Println(m)
	n := m["name"]
	fmt.Println(n)
	delete(m,"name")
	v, contrains := m["name"]
	if contrains {
		fmt.Println("contains name", v)
	} else {
		fmt.Println("not contains name", v)
	}
}
