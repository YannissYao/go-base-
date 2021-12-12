package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3

	//算数运算符
	//fmt.Println(a + b*c)
	//c++ //c=c+1
	//fmt.Println(c)
	//c-- //c=c-1
	//fmt.Println(c)
	////取余判断数字是否可以被整除
	//fmt.Println(b % c)

	//比较运算符
	//fmt.Println(a>b)
	//fmt.Println(b!=c)

	//逻辑运算符
	//fmt.Println(!(a > b))
	//fmt.Println(!(b!=c))

	//fmt.Println(a > b && b != c) //判断出左侧为false时,右侧不用判断
	fmt.Println(a > b || b != c)

	//赋值运算符
	//a += b
	//fmt.Println(a, b)
	//c *= b
	//fmt.Println(c, b)
}
