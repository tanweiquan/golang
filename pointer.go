package go_pointer

import (
	"flag"
	"fmt"
)

// 指针
// 指针是一种类型
// GO中的指针不能进行指针运算

// Go语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式
// 记住*Type是类型，表示指针类型，其中Type即可以是基础类型，也可以是自定义类型。
// &x为x的内存地址，接收内存地址的变量是指针，该指针的指针类型为*Type，这里Type为x的类型
/*
指针在GO中的两个重要的概念：
1、类型指针：允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
2、切片，由指向起始元素的原始指针、元素数量和容量组成。
*/
// 一个指针变量可以指向任何一个值的内存地址，它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节
// *Type表示一种指针类型。指针p的指针类型就是的*变量类型，如：当变量a是int类型，而p:=&a，则p为引用类型，它是指针变量，指向变量a的内存地址
// 值为内存地址的变量是指针，因此有声明指针：p:=&a 或 var p *Type = &a 或 p:=new(Type)
func main() {
	var a int = 13
	var b bool = false
	// 创建指针，法1(使用取值符&)
	// 指针p的类型为*int，且指针接收的值为a的内存地址
	p := &a
	fmt.Printf("p的类型为%T\n", p)
	// 指针ptr的类型为*bool，且指针接收的值为b的内存地址
	ptr := &b
	fmt.Printf("ptr的类型为%T\n", ptr)

	// 创建指针，法2(使用new()函数)
	pn := new(string) // *string类型指针

	// 用*符取内存地址对应的值
	va := *p
	vb := *ptr
	fmt.Printf("%v的类型为%T %v的类型为%T\n", va, va, vb, vb)
	*pn = "hello_world" // 根据*符赋值给指针
	fmt.Printf("%v的类型为%T\n", pn, pn)
	// 通过指针修改值
	x, y := 1, 2
	px := &x
	py := &y
	func(a, b *int) {
		*b, *a = *a, *b
	}(px, py)
	fmt.Println(x, y)

	// 使用flag包解析命令行后，用指针获取命令行参数
	m := flag.String("m", "hello", "ok") // 定义命令行参数,返回一个*string类型的指针
	flag.Parse()                         // 解析命令行
	vm := *m
	fmt.Println(vm)

	// 使用指向源变量的指针，来修改指针所指向的源变量的值
	type hhh struct {
		a string
		b int
	}
	aaa := hhh{ // 创建一个名为aaa的对象（值传递，即对对象的一次拷贝）
		a: "hello",
		b: 16,
	}
	xx := aaa        // 值拷贝，修改xx不影响aaa
	xx.b = 18        // 只改变了xx里的b，不影响aaa的b
	yy := &aaa       // 定义一个指向aaa的指针yy，修改yy的值会影响aaa
	(*yy).a = "bye"  // 语法糖为：yy.a="bye"
	fmt.Println(aaa) // {bye 16}

}
