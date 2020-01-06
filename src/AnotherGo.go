package main

import "fmt"

func main(){

	print("print")
	print("\n")
	print("%T 打印完成类型信息")
	print("%% 字面%")
	print("%V 格式")
	print("%+V 结构体会加字段名")
	print("%b 二进制表示")
	print("%c unicode码点表示的字符")
	print("%d 十进制表示")
	println("%s string []byte字符表示")
	println("")
	fmt.Println("%T","asd123__" )
	fmt.Println("%#v","asd123__" )
	fmt.Println("%+v","asd123__" )
	fmt.Println("%%")
	//todo panic defer
}
