package main

import "fmt"

func main() {
	//for i := 0; i < 5; i++ {
	//	if i == 2 || i == 4 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//双重for循环的continue
	//abc:for i := 0; i < 3; i++ {
	//	for j := 0; j < 3; j++ {
	//		if j == 1 {
	//			continue abc
	//		}
	//		fmt.Println(i, j)
	//	}
	//}

	//break在单个for循环中
	//for i := 0; i < 5; i++ {
	//	if i == 2 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	//break在双重for循环中
	myfor:for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j==1{
				break myfor
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("程序结束")
}
