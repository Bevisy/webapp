package main

import "fmt"

func main() {
	n := make(chan int)
	s := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			n <- i
		}
		close(n)
	}()

	go func() {
		for i := range n {
			s <- i * i
		}
		close(s)
	}()

	done := make(chan int) //同步 channel
	go func() {
		for i := range s {
			fmt.Println(i)
		}
		done <- 1
	}()

	<-done
}
