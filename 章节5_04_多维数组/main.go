package main

import "fmt"

func main() {
	//var arr [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	//arr := [3][3]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{7, 8, 9},
	//}
	//
	//fmt.Println(arr)
	//arr0:=arr[0]
	//fmt.Println(arr0[0],arr0[1],arr0[2])
	//fmt.Println(arr[0][0],arr[0][1],arr[0][2])

	arr := [2][2][2]int{
		{
			{
				1, 2,
			},
			{
				3, 4,
			},
		},
		{
			{
				3, 4,
			},
			{
				5, 6,
			},
		},
	}
	fmt.Println(arr[1][0][0])
}
