package main

import "fmt"

func main() {
	var name, age string
	fmt.Println("请输入您的年龄和姓名")
	//fmt.Scanln(&age, &name)
	//fmt.Scanf("%s\n%s",&name,&age)
	fmt.Scanf("%s%s", &name, &age)
	fmt.Println("输入的内容为:", name, age)
}
