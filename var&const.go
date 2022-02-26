package var_const

import "fmt"

// 变量
// 声明全局变量并赋值
var a int = 1

// 声明全局变量(类型推导)
var b = 3

// 一次声明多个变量
var m, n = "hello_world", 12

// 也可以这样一次声明多个变量
var (
	x = "你好"
	y = 66
)

// 程序运行期间不会改变的常量，一旦定义就不能修改。且常量声明后可以不使用。
// 定义全局常量
const pi = 3.14
const (
	ai = 1
	bi = 2
)

// 特别的有：
const (
	ci = 6
	xi // 6
	yi // 6
)

// 常量计数器iota,声明第一行常量时，iota=0，每新增一行常量，iota就自增1
const (
	av = iota //0
	bv = iota //1
	cv = iota //2
	nv = 100  //100，此时iota计数变为3
	mv = iota //4
	_  = iota //5，匿名常量，赋值但丢弃不使用
	xv = iota //6
	ov        //7
)

//多个常量声明在一行
const (
	a1, a2 = iota + 1, iota + 2 //a1:1,a2:2
	b1, b2 = iota + 1, iota + 2 //b1:2,b2:3
)

//定义数量级,符号<<表示左移,<<n表示左移n位,<<n=*(2^n)
const (
	_  = iota             //0
	KB = 1 << (10 * iota) //1<<(10*iota)=1*（2^(10*1))=2^10=1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 简短声明变量,只能在局部变量中使用
	c := 5
	fmt.Println(a, b, c)
	// 匿名变量
	u, _ := 2, 3
	fmt.Println(u)
	// 重新给a赋值
	a = 2
	fmt.Println(a)
	fmt.Println(m, n)
	fmt.Println(x, y)

	// 在函数中声明常量
	const k = 18
}
