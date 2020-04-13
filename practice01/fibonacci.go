package main

import "fmt"

func main() {
	var n = 10
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
