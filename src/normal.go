package main

import fmt2 "fmt"//可以取别名 / . "fmt" 可以省略包名 / _"fmt" 调用init
/**
 	程序初始化入口是main包 如果引入其他包递归去加载（链接）
	反向初始化/init

	非main包会生成一个 .a 文件在零食文件下，用于后续程序链接
	import 多个包要明确
 */
func init(){
	//别名
	fmt2.Println("如果有init函数，先执行init")
}

func main(){
	fmt2.Print("init initialization")

}
