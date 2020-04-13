package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()

	//Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	done := make(chan int)
	//Printer
	go func() {
		fmt.Println(<-squares)
		done <- 1
	}()

	<-done
}
