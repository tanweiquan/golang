package values_pointer

import "fmt"

func main() {
	// a,b 是一个值
	a := 5
	b := 6

	fmt.Println("a的值：", a) // a的值： 5

	// 指针变量 c 存储的是变量 a 的内存地址
	c := &a
	fmt.Println("a的内存地址：", c) // a的内存地址： 0xc000128058

	// 指针变量不允许直接赋值，需要使用 * 获取引用
	//c = 4

	// 将指针变量 c 指向的内存里面的值设置为4
	*c = 4
	fmt.Println("a的值：", a) // a的值： 4

	// 指针变量 c 现在存储的是变量 b 的内存地址
	c = &b
	fmt.Println("b的内存地址：", c) // b的内存地址： 0xc000128070

	// 将指针变量 c 指向的内存里面的值设置为4
	*c = 8
	fmt.Println("a的值：", a) // a的值： 4
	fmt.Println("b的值：", b) // b的值： 8

	// 把指针变量 c 赋予 c1, c1 是一个引用变量，存的只是指针地址，他们两个现在是独立的了
	c1 := c
	fmt.Println("c的内存地址：", c)   // c的内存地址： 0xc000128070
	fmt.Println("c1的内存地址：", c1) // c1的内存地址： 0xc000128070

	// 将指针变量 c 指向的内存里面的值设置为4
	*c = 9
	fmt.Println("c指向的内存地址的值", *c)   // c指向的内存地址的值 9
	fmt.Println("c1指向的内存地址的值", *c1) // c1指向的内存地址的值 9

	// 指针变量 c 现在存储的是变量 a 的内存地址，但 c1 还是不变
	c = &a
	fmt.Println("c的内存地址：", c)   // c的内存地址： 0xc000128058
	fmt.Println("c1的内存地址：", c1) // c1的内存地址： 0xc000128070
}
