package main

/**
闭包就是匿名函数

 */
func main(){

	f:= func() {
		println("i am a closure")
	}
	f()

	/**
		引用上下文变量
	 */
	count  :=1
	ff:= func() {
		count++
		println(count)
	}
	ff()
	ff()
	ff()
}
