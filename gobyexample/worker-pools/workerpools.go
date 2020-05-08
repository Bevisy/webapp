package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	//预备两个通道，模拟任务发布和处理
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//启动三个worker，准备处理jobs，初始处于阻塞状态
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//分发5份任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//获取任务处理结果，等待全部worker处理完毕
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// time go run workerpools.go
	//运行程序，显示 5 个任务被多个 worker 执行。
	//尽管所有的工作总共要花费 5 秒钟，但该程序只花了 2 秒钟， 因为 3 个 worker 是并行的。
}
