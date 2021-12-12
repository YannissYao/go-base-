package main

import "fmt"

func main() {
	var arr1 = [3]int{}
	arr2 := [3]int{1, 2, 3}
	arr3 := [8]int{1, 2, 3}
	fmt.Println(arr1, arr2, arr3)
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4)
	fmt.Println(len(arr1), len(arr2), len(arr3), len(arr4))
	//数组赋值和查看元素内容
	arr1[1] = 6
	fmt.Println(arr1, arr1[1], arr1[2])

	arr5 := arr1
	fmt.Println(arr1, arr5)
	fmt.Printf("%p %p", &arr1, &arr5)
	arr5[1] = 7
	fmt.Println(arr1, arr5)
}
