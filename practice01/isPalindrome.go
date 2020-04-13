package main

import "fmt"

func main() {
	var n int = 1
	fmt.Println(isPalindrome(n))
}

func isPalindrome(x int) bool {
	var num int
	var tmp int = x
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		for x != 0 {
			num = num*10 + x%10
			x /= 10
		}
		if num == tmp {
			return true
		} else {
			return false
		}
	}
}
