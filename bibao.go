package main

import "fmt"

// 闭包(闭包=函数+外部变量的引用),闭包就是能够读取其他函数内部变量的函数。
// 闭包是一个函数，这个函数包含了他外部作用域的变量
/*
底层的原理
1.函数可以作为返回值
2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找
*/
// 闭包的两个定义：
// 第一：必须要有在函数外部定义，但在函数内部引用的“自由变量”
// 第二：脱离了形成闭包的上下文，闭包也能照常使用这些自由变量

// 例：
// 闭包，且func(int)int作为返回值
func create() func(int) int {
	var x int = 1 // x在匿名函数外部定义，但在下面的匿名函数内部引用
	return func(i int) int {
		x += i
		return x
	}
}

//
func main() {
	// 将闭包的返回值函数赋值给fb变量
	fb1 := create() // 相当于fb:=func(int)int，但改函数变量的作用域在create里
	b1 := fb1(1)    // 调用闭包返回值函数
	fmt.Println(b1)
	fb2 := create() // 使用f2接收返回值函数
	b2 := fb2(1)    // f2调用后return的x为2
	// 重复调用相同函数，栈帧中调用在同一个栈帧空间里
	bt := fb2(2) // 优先往外部相同函数f2找变量x,即使用f2返回的x
	fmt.Println(b2)
	fmt.Println(bt)
}
