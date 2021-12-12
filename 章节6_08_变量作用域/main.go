package main

import "fmt"

func main() {
}

/*
局部变量演示
 */
//func demo1(){
//	//i声明在函数的{}内部,所以整个函数内部都可以使用
//	i:=1
//	if i>0{
//		//j声明在if{}内部,所以j只能在if中使用
//		j:=2
//		fmt.Println(j,i)
//	}
//	//fmt.Println(i,j)
//	fmt.Println(i)//只能访问i
//}
//
//func demo2(){
//	fmt.Println(i)
//}

/*
全局变量
 */

 var (
 	name="smallming"
 	age = 17
 )
 func demo1(){
 	fmt.Println("名字:",name,age)
 }
 func demo2(){
 	fmt.Println("年龄,",age,name)
 }

