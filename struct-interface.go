package main

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

// 多态：使用interface来实现 (对象类型不同时，调用相同方法，可能执行不同函数)
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
type American struct { //定义对象类型American
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

type Russian struct { //定义对象类型Russian
	name     string
	Language string
	food     string
}

func (a *Russian) Speak() {
	fmt.Println("俄罗斯人" + "说" + "要打美国人")
}
func (a *Russian) Eat() {
	fmt.Println("俄罗斯人" + "吃" + "什么？")
}
func main() {
	// 这里就看出了多态的，不同的对象类型，调用相同方法，执行的是不同函数
	a := &American{
		"jieli",
		"English",
		"汉堡",
	}
	r := &Russian{
		"达瓦",
		"Russian",
		"红肠",
	}
	a.Speak()
	a.Eat()
	r.Speak()
	r.Eat()
}
