package recursion

//递归: 函数调用自身，缩小范围
//fact()在到达fact(0)前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
