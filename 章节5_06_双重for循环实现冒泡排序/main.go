package main

import "fmt"

func main() {
	//for i := 0; i < 2; i++ {
	//	for j := 0; j < 2; j++ {
	//		fmt.Println(i, j)
	//	}
	//}

	//冒泡排序
	arr := [5]int{1, 7, 5, 6, 2}
	//第一次把7放到最后
	//[1,5,6,2,7]
	//for i:=0 ;i<len(arr)-1;i++{
	//	if arr[i]>arr[i+1]{
	//		arr[i],arr[i+1]=arr[i+1],arr[i]
	//	}
	//}
	////把6放到最后
	////[1,5,2,6,7]
	//for i:=0 ;i<len(arr)-1-1;i++{
	//	if arr[i]>arr[i+1]{
	//		arr[i],arr[i+1]=arr[i+1],arr[i]
	//	}
	//}
	////把5放到后面
	////[1,2,5,6,7]
	//for i:=0 ;i<len(arr)-1-2;i++{
	//	if arr[i]>arr[i+1]{
	//		arr[i],arr[i+1]=arr[i+1],arr[i]
	//	}
	//}
	////把2放到后面
	////[1,2,5,6,7]
	//for i:=0 ;i<len(arr)-1-3;i++{
	//	if arr[i]>arr[i+1]{
	//		arr[i],arr[i+1]=arr[i+1],arr[i]
	//	}
	//}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
