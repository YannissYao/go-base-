package main

import "fmt"

func main() {
	//a:=1.5
	//fmt.Println(&a)
	//a=2
	//fmt.Println(&a)
	//fmt.Println(a)
	//
	//b:=a
	//b=3
	//fmt.Println(&b,&a)
	//fmt.Println(b,a)
	//
	//
	//c:=&a
	//fmt.Printf("%T",c)

	var a *int
	fmt.Println(a)
	fmt.Println(a == nil)

	b := 123
	a = &b

	fmt.Println(a, &b)

	*a = 456
	*&b = 789
	fmt.Println(*a, b)

}
