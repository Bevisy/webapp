package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` { // 反引号 用来创建原生的字符串字面量，这些字符串可能由多行组成(不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML以及正则表达式
			fmt.Printf("\r%c", r) // \r 是回车符（return），作用是使光标移动到本行的开始位置；
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	go spinner(100 * time.Millisecond) //新建goroutine，创建完则返回，程序接着往下走;当主函数返回时，所有的goroutine都会直接打断，程序退出
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN) // 双引号用来创建可解析的字符串字面量(支持转义，但不能用来引用多行)
}
