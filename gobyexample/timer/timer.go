package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器表示在未来某一时刻的独立事件。
	//告诉定时器一个需要等待的时间，然后定时器提供一个用于通知的通道。
	timer1 := time.NewTimer(2 * time.Second)

	//<-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	//如果你需要的仅仅是单纯的等待，使用 time.Sleep 就够了。
	//使用定时器的原因之一就是，你可以在定时器触发之前将其取消。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
