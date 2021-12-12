package main

import "fmt"

//func main() {
//	ch1 := make(chan int, 1)
//	ch2 := make(chan string, 1)
//	//ch1 <- 1
//	//ch2 <- "smallming"
//	select {
//	case a1 := <-ch1:
//		fmt.Println(a1)
//	case a2 := <-ch2:
//		fmt.Println(a2)
//	//default:
//	//	fmt.Println("default")
//	}
//}

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
		}

	}

}
