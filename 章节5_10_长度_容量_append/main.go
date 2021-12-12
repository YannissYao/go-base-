package main

import "fmt"

func main() {
	//不指定切片容量
	//s := make([]int,0)
	//fmt.Println(s==nil)
	//fmt.Printf("%p\n",s)
	//fmt.Println(len(s),cap(s))
	////指定切片的容量
	//s1 := make([]int,0,3)
	//fmt.Println(len(s1),cap(s1))


	//append添加一个或多个值时,长度和容量的变化
	//s := make([]string,0)
	//fmt.Println(len(s),cap(s))
	//s=append(s,"smallming")
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))
	//s= append(s,"老张")
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))
	//s= append(s,"佳明哥")
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))
	//s=append(s,"4","5","6","7","8","9")
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))
	//s=append(s,"10")
	//fmt.Println(s)
	//fmt.Println(len(s),cap(s))

	s:= make([]int,0)
	s1:=[]int{1,2,3,4}
	s=append(s,s1...)
	fmt.Println(s)
	fmt.Println(len(s),cap(s))
}
