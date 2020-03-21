package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name",orm:"name",ini:"name"`
	Addr string `json:"addr"`
}

func main() {
	p:=Person{
		Name: "李广昌",
		Addr: "北京",
	}
	//序列化
	marshal, _ := json.Marshal(p)
	fmt.Println(string(marshal))
	//反序列化
	str:=`{"name":"lgc","addr":"北京"}`
	var p1 Person
	json.Unmarshal([]byte(str), &p1)
	fmt.Printf("%#v\n",p1)
}
