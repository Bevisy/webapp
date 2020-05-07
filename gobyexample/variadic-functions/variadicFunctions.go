package main

import "fmt"

//变参函数
func Sum(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3))
	nums := []int{1, 2, 3, 4}
	fmt.Println(Sum(nums...))
}
