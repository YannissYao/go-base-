package main

import (
	"fmt"
)

//函数变量是引用类型
//func main() {
//	var a func()
//	a = b
//	fmt.Printf("%p %p",a,b)
//}
//
//func b() {
//	fmt.Println("b")
//}

/*
函数作为参数
 */
//func main() {
//	mydo(func(name string) {
//		fmt.Println(name)
//	})
//}
//
//func mydo(arg func(name string)) {
//	fmt.Println("执行mydo")
//	arg("smallming")
//}

/*
函数作为返回值
 */
func main() {
	result:=a()
	r2:=result()
	fmt.Println(r2)
}
func a() func() int{
	return func() int {
		return 110
	}
}
