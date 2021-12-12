package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("11111")
	//time.Sleep(2e9)
	//fmt.Println("22222")
	//fmt.Println("3333")

	fmt.Println("程序开始")
	time.AfterFunc(3e9, func() {
		fmt.Println("延迟执行")
	})
	time.Sleep(4e9)
	fmt.Println("程序结束")
}
