package _struct

import (
	"fmt"
	"testing"
)

type mySt struct {
	id   string
	name string
}

func Test_Struct(t *testing.T) {

	//definition struct
	type mySt struct {
		id   string
		name string
		sex  string
		tel  string
	}
	//init struct
	s := mySt{"1", "李广昌", "男", "123456"}
	fmt.Printf("out myStruct %v\n", s)
	fmt.Printf("out myStruct.name %v\n", s.name)

	//definition around struct

	type AroundStruct struct {
		structName string
		myStruct   mySt
	}
	/**
	打印结构体用 %+v %v不会打印字段名
	*/
	fmt.Printf("out my around struct %+v\n", AroundStruct{"我的嵌套结构体", mySt{"1", "走着走着就被嵌套来了", "3", "3"}})
	fmt.Printf("out my around struct %v\n", AroundStruct{"我的嵌套结构体", mySt{"1", "走着走着就被嵌套来了", "3", "3"}})

	/**
	结构体指针 结构体指针可以直接获取属性值
	*/
	sp := &mySt{"1", "1", "1", "1"}
	fmt.Printf("out struct point value %s \n", sp.name)
}

func Test_StructMethod(t *testing.T) {

	m := mySt{"1", "2"}
	m.service("lgc")

}
/**
	结构体方法  结构体 拥有方法
 */
func (m *mySt) service(param string) {
	fmt.Printf("input param is %s\n", param)
	m.name = param
	fmt.Printf("out struct %+v\n",m)
}
