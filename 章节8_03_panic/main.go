package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer的内容")
	fmt.Println("111111111")
	panic("出现panic了,呜呜呜,剩余代码不能执行")
	//os.Exit(0)
	fmt.Println("2222222223")
}
