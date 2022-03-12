package go_struct

import "fmt"

/*
Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
*/
// 结构体的定义
/*
type struct_variable_type struct {
	member definition;
	member definition;
	...
	member definition;
}
*/
// 例子
type person struct {
	name string
	age  int
}

func main() {
	// 初始化结构体
	// 1.按照顺序提供初始化值
	p1 := person{"Tom", 25}
	// 2.通过field:value的方式初始化，这样可以任意顺序
	p2 := person{age: 24, name: "Tom"}
	// 3.new方式,未设置初始值的，会赋予类型的默认初始值
	p3 := new(person)
	p3.age = 24
	fmt.Println(p1, p2, p3)
}
