package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //创建 WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //每启动一个新的goroutine，计数加1
		go func(i int) {
			fmt.Printf("start task %d\n", i)
			defer wg.Done() //goroutine退出，则减一
		}(i)
	}
	wg.Wait()

	fmt.Println("Finished.")
}
