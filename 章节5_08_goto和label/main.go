package main

import "fmt"

func main() {
	//goto结合if演示
	//	i := 6
	//	if i == 6 {
	//		goto Loop
	//	}
	//	fmt.Println("执行完if")
	//Loop:
	//	fmt.Println("loop")
	//	if i!=6{
	//		goto Loop1
	//	}
	//Loop1:
	//	fmt.Println("Loop1")
	//	fmt.Println("程序结束")

	//结合for循环演示
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			goto abc
		}
	}
	fmt.Println("执行完for")
abc:
	fmt.Println("abc")
}
