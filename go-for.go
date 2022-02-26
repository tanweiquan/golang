package go_for

import "fmt"

// for（循环结构）
// 语句 1 在循环（代码块）开始前执行 (i:=?) 且变量的作用域:只在循环语句中生效
// 语句 2 定义运行循环（代码块）的条件 (i<?)
// 语句 3 在循环（代码块）已被执行之后执行 (i++)

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种1
	var g = 5
	for ; g < 10; g++ {
		fmt.Println(g)
	}
	// 变种2
	var h = 11
	for h < 15 {
		fmt.Println(h)
		h++
	}
	// 变种3
	a := 1
	for {
		a++
		if a > 10 {
			fmt.Println("循环退出")
			break // 退出循环的方式有两种：break和goto lable
		} else {
			continue // 继续循环
		}
	}
	// 遍历
	// for range(键值循环)遍历数组，切片，字符串,map及通道（channel）。
	// 根据目录找内容
	// 1.数组，切片，字符串返回索引和值；
	// 2.map返回键和值；
	// 3.通道（channel）只返回通道内的值

	// 例：
	str := "hello沙河"
	for n, m := range str {
		fmt.Printf("%d %c\n", n, m) // 这里的%c根据码值格式化输出为对应的字符
	}

}
