package main

import "fmt"

const (
	//globalArr = [...]int{2, 3, 4}
	/**
	常量只能是 字符串 布尔 整数
	 */
)

func init() {
	numbers := [3]int{1, 2, 3}
	number := [...]int{1, 2, 3, 4}
	fmt.Printf("numbers out : '%v'", numbers)
	println()
	fmt.Printf("number out : '%v'", number)
	println()
}

/**
range 会返回 数组索引和对应的值
*/
func sum(arr []int) (all int) {
	for _, number := range arr {
		all += number
	}
	return
}

func main() {
	arr:=[]int{5,6,7}
	println(sum([]int{1, 2}))
	/**
		数组不方便 引入了 slice
	 */
	fmt.Printf("arr length is %d\n",len(arr))
	var s  =arr[0:1]
	fmt.Printf("slice length is %d\n",len(s))
	sli:=[]int{1,2,3}
	fmt.Printf("slice capacity is %d\n",cap(sli))
	//make 创建切片
	a:=make([]int,10)
	fmt.Printf("make(type,size) make a slice out  cap '%d', len '%d'\n",cap(a),len(a))
	ints := append(a, 1)
	fmt.Printf(" use append add  element %v\n",ints)
	for index,value :=range ints  {
		fmt.Printf("range out slice : key:%d  value:%d \n",index,value)
	}
	for _,value :=range ints  {
		fmt.Printf("ignore index : key:%d  value:%d \n",value)
	}
	for index,_ :=range ints  {
		fmt.Printf("ignore value : key:%d  value:%d \n",index)
	}
	for index :=range ints  {
		fmt.Printf("use one element default ignore value : key:%d  value:%d \n",index)
	}
	ns:=[]int{4:5}
	println(len(ns))
	for index,value:=range ns{
		fmt.Printf("ns index: %d,value: %d\n",index,value)
	}

}
