package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(1e9)
			fmt.Println("第", i, "次执行")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Print("程序结束")
}
