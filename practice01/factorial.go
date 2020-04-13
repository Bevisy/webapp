package main

import "fmt"

func main() {
	fmt.Printf("阶乘结果是：%d\n", factorial(10))
}
func factorial(n uint64) (ret uint64) {
	if n > 0 {
		ret = n * factorial(n-1)
		return ret
	}
	return 1
}
