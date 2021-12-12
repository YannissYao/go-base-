package main

import "fmt"

func main() {
	fmt.Println("执行main")
	//demo()
	a, b := 1, 2
	sum := add(a, b)
	fmt.Println(sum)
}

func demo() {
	fmt.Println("执行了demo")
}

func add(c, d int) (s int) {
	//fmt.Println(c+d)
	s= c+d
	return
}
