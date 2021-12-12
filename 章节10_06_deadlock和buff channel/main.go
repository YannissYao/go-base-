package main

import "fmt"

func main() {
	ch := make(chan int,3)
	ch <- 1
	<-ch
	ch <- 1
	ch <- 1
	ch <- 1
	fmt.Println("程序结束")
}