package main

import "fmt"

func main() {
	/*
	多if
	在执行时,无论前面的if是否成立,后面的if都会进行判断
	 */
	score := 65
	if score >= 60 {
		fmt.Println("及格")
	}
	if score < 60 {
		fmt.Println("不及格")
	}

	/*
	在if 的表达式中声明变量
	 */
	//if  score := 65;score >= 60 {
	//
	//	fmt.Println("及格",score)
	//}
}
