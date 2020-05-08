package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	//我们可以通过缓冲通道来完成此任务。 burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstyLimiter := make(chan time.Time, 3)

	//填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//第二批请求，由于爆发（burstable）速率控制，我们直接连续处理了 3 个请求， 然后以大约每 200ms 一次的速度，处理了剩余的 2 个请求。
	for req := range burstyRequests {
		<-burstyLimiter //只读，前三个直接并发启动
		fmt.Println("request", req, time.Now())
	}
}
