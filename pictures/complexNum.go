package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex64 = complex(3, 4) // 使用 complex 内建函数
	var y complex64 = complex(6, 8)
	//var z complex128 = complex(1, 2)
	fmt.Println(x)
	fmt.Printf("%v\n", y)
	// 在 go 里，复数可以直接做四则运行
	fmt.Println(x + y)
	fmt.Println(x * y)
	fmt.Println(x / y)

	fmt.Println()

	a := 1 + 2i // 也可以使用字面量。如果不指定类型，则推导类型是 complex128
	b := 2 + 3i
	c := cmplx.Sqrt(-4.41) // 对负数开根号，也可以得到复数。
	fmt.Println(a * b)
	fmt.Println(c)
	fmt.Println(real(a)) // 计算实部
	fmt.Println(imag(a)) // 计算虚部
}
