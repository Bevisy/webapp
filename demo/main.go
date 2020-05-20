package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // x 只会输出 0、2、4、6、8，为什么？ TODO
			fmt.Println(x)
		case ch <- i:
		}
	}
}

//倒计时
//func main() {
//	fmt.Println("Starting...")
//	t := time.Tick(time.Second * 1)
//	for i:= 10;i>0;i--{
//		<-t
//	}
//	fmt.Println("End.")
//}
//goroutine 泄漏；此处主goroutine会退出，泄漏不造成实际影响
//func main() {
//	ch := make(chan int)
//	go func() {ch<-1}()
//	go func() {ch<-1}()
//	go func() {ch<-1}()
//	<-ch
//}
//带缓冲的channel，查看长度和容量
//func main() {
//	ch := make(chan string, 3)
//	fmt.Println(len(ch), cap(ch)) //查看channel实际数据长度，和channel容量
//	ch <- "A"
//	fmt.Println(len(ch), cap(ch)) //查看channel实际数据长度，和channel容量
//}

//func main() {
//	nature := make(chan int)
//	square := make(chan int)
//
//	go func() {
//		for i := 1; i < 100; i++ {
//			nature <- i
//		}
//		close(nature)
//	}()
//
//	go func() {
//		//for x := range nature { //使用range遍历通道
//		//	square <- x * x
//		//}
//		for {
//			x, ok := <-nature //多接收一个结果，多接收 的第二个结果是一个布尔值ok，
//			// ture表示成功从channels接收到值，false表示channels已经被关闭并且 里面没有值可接收。
//			if !ok {
//				break
//			}
//			square <- x * x
//		}
//		close(square)
//	}()
//
//	for x := range square {
//		fmt.Printf("output: %d\n", x)
//	}
//
//}

//func spinner(delay time.Duration) {
//	for {
//		for _, s := range `\|/-` {
//			fmt.Printf("\r%c", s)
//			time.Sleep(delay)
//		}
//	}
//}
