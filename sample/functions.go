package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

//return multi values
func vals() (int, int) {
	return 3, 7
}

//Variadic Functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	v1, v2 := vals()
	fmt.Printf("return values: %d %d\n", v1, v2)

	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
