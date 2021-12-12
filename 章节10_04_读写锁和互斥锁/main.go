package main

import (
	"sync"
	"fmt"
)

//var (
//	num = 100
//	wg sync.WaitGroup
//	m sync.Mutex
//)
//func demo(){
//	m.Lock()
//	for i:=0;i<10;i++{
//		num = num -1
//	}
//	m.Unlock()
//	wg.Done()
//}
//
//func main() {
//	wg.Add(10)
//	for i :=0;i<10;i++{
//		go demo()
//	}
//	wg.Wait()
//	fmt.Println(num)
//	fmt.Println("程序结束")
//}

/*
测试读写锁
 */
func main() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(10)
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			rwm.Lock()
			m[j] = j
			fmt.Println(m)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("程序结束")
}
