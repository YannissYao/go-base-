package main

import "fmt"

func main() {
	//1. 先声明后赋值
	//var name string
	//name = "smallming"
	//fmt.Println(name)

	//2. 声明同时赋值
	//var name string = "smallming"
	//fmt.Println(name)

	//3. 声明同时赋值(省略类型)
	//var name = "smallming"
	//fmt.Println(name)
	//fmt.Printf("%T",name)

	//4. 短变量
	//name := "smallming"
	//fmt.Println(name)

	//5. 一次声明多个变量
	//var a,b,c int
	//a,b,c = 1,2,3
	//fmt.Println(a,b,c)

	//6.声明并赋值
	//var a,b,c = 1,2,false
	//fmt.Println(a,b,c)

	//7.使用var声明并赋值
	//var (
	//	a=1
	//	b=2
	//	c=false
	//)
	//fmt.Println(a,b,c)


	//8.使用短变量声明并赋值多个变量时要求必须至少有一个变量没有声明
	var (
		a=1
		b=2
		c=3
	)
	b,c,d:=4,5,6
	fmt.Println(a,b,c,d)
}
