package main

import "fmt"

func main() {
	//当变量取及格固定取值时
	//switch num := 18;num {
	//default:
	//	fmt.Println("没有这个进制")
	//case 2:
	//	fmt.Println("二进制")
	//case 8:
	//	fmt.Println("八进制")
	//case 10:
	//	fmt.Println("十进制")
	//case 16:
	//	fmt.Println("十六进制")
	//}

	/*
	条件为范围时
	 */
	//score := 99
	//switch {
	//case score >= 90:
	//	fmt.Println("优秀")
	//case score >= 80:
	//	fmt.Println("良好")
	//case score >= 70:
	//	fmt.Println("中等")
	//case score >= 60:
	//	fmt.Println("及格")
	//default:
	//	fmt.Println("不及格")
	//}

	/*
	case 后可以有多个值,每个值之间使用逗号分隔
	 */
	//month := 2
	//switch  month {
	//case 1, 3, 5, 7, 8, 10, 12:
	//	fmt.Println("31天")
	//case 2:
	//	fmt.Println("28/29天")
	//default:
	//	fmt.Println("30天")
	//}

	/*
	穿透
	 */
	switch num := 8; num {
	default:
		fmt.Println("default")
		break
		fmt.Println("default2")
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
		//fallthrough 不能在最后一个case/default中
	}
	fmt.Println("程序结束")
}
