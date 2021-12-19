package main

import (
	"fmt"
)

func main() {
	demo(1)
	var i interface{} = 1
	re, ok := i.(float64)
	fmt.Println(re, ok)
}

func demo(i interface{}) {
	_, ok := i.(int)
	if ok {
		fmt.Println("参数是int类型")
		return
	}
	_, ok = i.(float64)
	if ok {
		fmt.Println("参数是float64类型")
		return
	}
	fmt.Println("参数类型不确定,不是int或float64")
}
