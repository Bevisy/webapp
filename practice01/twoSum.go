package main

import "fmt"

func main() {
	n := []int{3, 2, 4}

	fmt.Println(twoSum(n, 6))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
