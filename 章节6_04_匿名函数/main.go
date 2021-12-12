package main

import "fmt"

func main() {

	//匿名函数-无参数无返回值
	func() {
		fmt.Println("无参数无返回值匿名函数")
	}()

	//有参数的匿名函数
	func(name string) {
		fmt.Println("名字为:", name)
	}("smallming")

	//有返回值的匿名函数
	name := func() string {
		fmt.Println("有返回值的匿名函数")
		return "佳明哥"
	}() //此处表示调用
	fmt.Println(name)
}
