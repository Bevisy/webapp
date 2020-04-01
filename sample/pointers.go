package main

import "fmt"

//局部变量修改，不影响原值
func zeroval(ival int) {
	ival = 0
}

//取地址赋值，原值被修改
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
