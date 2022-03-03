package go_defer

import "fmt"

// defer相当于弹夹一样，先进后出
// 下面执行顺序是 var a int->fmt.Println(a)->a=100->a++
func f1() {
	var a int
	defer func() {
		a++
	}()
	fmt.Println(a)
	defer func() {
		a = 100
	}()
}

// go语言中函数的return不是原子操作，在底层是分为两步来执行的
// 第一步：返回值赋值
// defer
// 第二步：真正的返回值

// 返回值未命名
// 下面执行顺序是：局部变量x=6 -> 返回值=x=6 -> 局部变量=7 -> 真正ret
func fs1() int {
	x := 6
	defer func() {
		x++
	}()
	return x
}

// 返回值已命名
// 下面执行顺序是：局部变量y=6 -> 返回值m=12 -> 局部变量y=7 -> 真正ret
func fs2() (m int) {
	y := 6
	defer func() {
		y++
	}()
	return 12
}

// 下面执行顺序是：局部变量x=6 -> 返回值y=x=6 -> 局部变量=x=7 -> 真正ret
func fs3() (y int) {
	x := 6
	defer func() {
		x++
	}()
	return x
}

// 下面执行顺序是：局部变量x=6 -> 返回值y=x=6 -> 局部变量=x=7 -> 返回值y=y+1=7 -> 真正ret
func fs4() (y int) {
	x := 6
	defer func() {
		y++
	}()
	return x
}

// 进阶
// 下面执行顺序是：局部变量=0 -> 返回值x=5 -> 局部变量=x=6 -> 真正ret
func fss() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	f1()
	a := fs1()
	fmt.Println(a)
	b := fs2()
	fmt.Println(b)
	c := fs3()
	fmt.Println(c)
	d := fss()
	fmt.Println(d)
}
