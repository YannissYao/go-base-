package main

import "fmt"

func main() {
	//调用不接收
	show()

	//接收
	a, b, c := show()
	fmt.Println(a, b, c)

	e, f, _ := show()
	fmt.Println(e, f)
}

func show() (name string, age int, addr string) {
	fmt.Println("执行了show")
	name = "smallming"
	age = 17
	addr = "北京海淀"
	return
}
