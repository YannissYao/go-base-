package main

import "fmt"

func main() {
	//f是返回值函数
	f := closure()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	f1 := closure()
	fmt.Println(f1())
	fmt.Println(f1())

	fmt.Println(f())

	fmt.Printf("%p %p", f, f1)
}

func closure() func() int {
	i := 1
	return func() int {
		//fmt.Printf("%p\n",&i)
		i = i + 1
		return i
	}
}
