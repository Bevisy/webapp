package main

import (
	"sync"
)

//限制goroutine启动数量，并在主goroutine结束前返回
//func main() {
//
//	//方法一
//	var wg sync.WaitGroup   //控制主goroutine退出
//	ch := make(chan int, 5) //限制goroutine数量
//	for i := 0; i < 10; i++ {
//		ch <- 1 //channel满时，阻塞新的goroutine启动
//		wg.Add(1)
//		go func(i int) {
//			fmt.Printf("start tasjk %d\n", i)
//			time.Sleep(time.Second * 1)
//			<-ch //子goroutine运行结束后，读取通道，消除一次阻塞
//			defer wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("Finished.")
//}

//方法二

type Pool struct {
	queue chan int
	wg *sync.WaitGroup
}

func NewPool(size int) *Pool{
	if size <= 0 {
		size = 1
	}

	return &Pool{
		queue: make(chan int, size),
		wg: &sync.WaitGroup{}
	}
}

func (p *Pool)Add(delta int) {

}