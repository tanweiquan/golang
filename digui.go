package digui

import "fmt"

/*
Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列，链表倒置等。
*/

// 计算阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 { // 设置退出条件
		result = n * Factorial(n-1) // 递归
		return result
	}
	return 1
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}
