package main

import "fmt"

//根据数组的特征，依次判断相邻的两个元素是否相同，相同则跳过，不同则取值保存，直至数组遍历结束
//循环只有一层所以算法时间复杂度为o(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ret int
	for i := range nums {
		if nums[ret] == nums[i] {
			continue
		} else {
			ret += 1
			nums[ret] = nums[i]
		}
	}
	return ret + 1
}

func main() {
	nums := []int{0, 1, 1, 2, 2, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
