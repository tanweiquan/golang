package struct_interface

import "fmt"

// 封装：首字母大小写方式代表公有私有权限
/*
type 类型名 struct{
    字段名 字段类型
    字段名 字段类型
    ...
}
*/
type Person struct {
	name string
}

func (p *Person) Walk() {
	fmt.Println(p.name)
}

// 继承：使用内嵌的方式，对结构体struct进行组合
type Chinese struct {
	p    Person
	skin string
}

func (c *Chinese) GetSkin() string {
	return "黄皮肤"
}

// 多态：使用interface来实现(可以通过对象访问对象里的字段，同时可以调用对象绑定的方法)
// 关联interface的结构体(类)要实现interface里的所有方法
/*
type 接口名 interface{
	方法名1(参数1，参数2...)(返回值1，返回值2...)
	方法名2(参数1，参数2...)(返回值1，返回值2...)
	...
}
*/
type Human interface {
	Speak()
	Eat()
}
type American struct {
	name     string
	Language string
	food     string
}

func (a *American) Speak() {
	fmt.Println("美国人" + a.name + "说" + a.Language)
}
func (a *American) Eat() {
	fmt.Println("美国人" + a.name + "吃" + a.food)
}
func main() {
	a := &American{
		"jieli",
		"English",
		"汉堡",
	}
	a.Speak()
	a.Eat()
	fmt.Println(a.name)
	fmt.Println(a.Language)
	fmt.Println(a.food)
}
