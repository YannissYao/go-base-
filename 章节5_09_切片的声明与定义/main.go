package main

import "fmt"

func main() {
	//声明时为空
	//var slice []string
	//fmt.Println(slice)
	//fmt.Println(slice==nil)
	//fmt.Printf("%p",slice)


	//给切片赋初值
	//names:=[]string{"smallming","佳明哥"}
	//fmt.Println(names)
	//fmt.Println(names[0],names[1])
	//names[0]="老张"
	//fmt.Println(names)


	//切片是引用类型
	names:=[]string{"smallming","佳明哥"}
	fmt.Println(names)
	names1:=names
	fmt.Println(names1,names)
	names1[0]="老张"
	fmt.Println(names1,names)
	fmt.Printf("%p %p",names1,names)
}
