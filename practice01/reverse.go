package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(-123))

}

func reverse(x int) int {
	var num int
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	//fmt.Printf("%b,%b\n", MaxInt32, MinInt32)
	if num > MaxInt32 || num < MinInt32 {
		return 0
	}
	return num
}
