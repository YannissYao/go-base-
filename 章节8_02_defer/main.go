package main

import "fmt"

func main() {
	/*
	单个defer
	 */
	//fmt.Println("打開連接")
	////defer fmt.Println("关闭连接")
	//defer func() {
	//	fmt.Println("关闭连接")
	//}()
	//fmt.Println("执行功能")
	/*
	多个defer
	 */
	 //fmt.Println("打开A")
	 //defer fmt.Println("关闭A")
	 //fmt.Println("打开B")
	 //defer fmt.Println("关闭B")
	fmt.Println(demo())
}

func demo() (z int){
	i:=0
	defer func() {
		z=i+2
	}()
	return i
}