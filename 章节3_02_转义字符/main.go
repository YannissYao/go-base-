package main

import "fmt"

func main() {
	//输出结果
	//fmt.Printf("%d %x %X %o %b",18,11,12,18,18)

	//把转换后的结果获取到
	//a := fmt.Sprintf("%x", 18)
	//fmt.Println(a)

	//浮点数
	//fmt.Printf("%f",1.5)

	//布尔类型
	//fmt.Printf("%t", false)

	//字符
	//fmt.Printf("%c",65)

	//输出字符串
	//fmt.Printf("%s %q","你好","你好")

	//内置格式
	//fmt.Printf("%v", "你好")

	//类型
	//fmt.Printf("%T","false")

	//地址
	i := 5
	fmt.Printf("%d", &i)

	//百分号
	//fmt.Printf("花费了%d%%总钱数",20)

	//换行和缩进
	fmt.Printf("你好吗?\n我很\n好")
}
