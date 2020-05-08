package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: //读取通道，如果通道无值，则等待
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped.")
}
