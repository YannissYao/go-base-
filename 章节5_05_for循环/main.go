package main

import "fmt"

func main() {
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//死循环
	//for ; ; {
	//	fmt.Println("执行")
	//}

	//i:=0
	//for ;i<5;{
	//	fmt.Println(i)
	//	i++
	//}
	//
	//j:=0
	//for j<5{
	//	fmt.Println(j)
	//	j++
	//}

	//遍历数组(一)
	arr := [3]string {"张","佳明哥","smallming"}
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	//遍历数组,结合range
	//i 脚标
	//n 元素内容 n=arr[i]
	for _,n :=range arr{
		fmt.Println(n)
	}
}
