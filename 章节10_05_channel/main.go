package main

import (
	"fmt"
)

/*
同步
主协程和子协程之间的通信
 */
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		fmt.Println("执行")
//		fmt.Println("结束")
//		ch <- 998
//	}()
//	a := <-ch
//	fmt.Println(a)
//	fmt.Println("程序结束")
//}

/*
两个子协程的通信
 */

//func main() {
//	ch1 := make(chan string)
//
//	ch2 :=make(chan int)
//
//	go func() {
//		ch1 <- "传送给另一个协程的数据"
//		ch2 <- 1
//	}()
//	go func() {
//		content := <- ch1
//		fmt.Println("取出数据成功:", content)
//		ch2<-2
//	}()
//
//	<-ch2
//	<-ch2
//	fmt.Println("程序结束")
//}

/*
使用for range取出channel中数据
 */

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func() {
		for i := 97; i <= 97+26; i++ {
			ch1 <- fmt.Sprintf("%c", i)
		}
		ch2 <- 1
	}()

	go func() {
		for n := range ch1 {
			fmt.Println(n)
		}
		ch2<-2
	}()

	<-ch2
}
