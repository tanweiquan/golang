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
// 结构体名大写导出，结构体小写不可以导出
type person struct {
	Name string // 大写导出成员
	age  int    // 小写不可以导出
}

func main() {
	// 初始化结构体
	// 1.按照顺序提供初始化值
	p1 := person{"Tom", 25}
	// 2.通过field:value的方式初始化，这样可以任意顺序
	p2 := person{age: 24, Name: "Tom"}

	// 使用&或者new来创建一个引用类型的结构体(即创建指针)
	p4 := &person{"hellokitty", 10}
	// 4.new方式,未设置初始值的，会赋予类型的默认初始值
	p3 := new(person)
	p3.age = 24
	fmt.Println(p1, p2, p3, p4)
	/*
		引用类型的变量，传入函数时，虽然也是传值，但拷贝的是引用类型的内存地址，可以说拷贝了一个引用，这个引用指向了函数体外的某个结构体，使用这个引用在函数里修改结构体的值，外面函数也会发现。
		如果传入的不是引用类型的结构体，而是值类型的结构体，那么会完整拷贝一份结构体，该结构体和原来的结构体就没有关系了
		内置的数据类型切片 slice 和字典 map 都是引用类型，不需要任何额外操作，所以传递这两种类型作为函数参数，是比较危险的，开发的时候需要谨慎操作
	*/
}
