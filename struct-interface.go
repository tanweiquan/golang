package main

import "fmt"

// 封装：对外公开接口，增强安全，简化编程。（首字母大小写方式代表公有私有权限）
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

// 继承：子类继承父类，自动拥有父类的属性和方法。（使用内嵌的方式，对结构体struct进行组合）
type Chinese struct {
	p    Person
	skin string
}

func (c *Chinese) GetSkin() string {
	return "黄皮肤"
}

// 多态：使用interface来实现 (对象类型不同时，调用相同方法，可能实现不同的功能)
// 实现接口interface：将对象类型不同而能实现不同功能的同名方法，放在接口interface里；接口名自定义即可，不影响程序。
// 当子类继承父类时，如果子类绑定的方法跟父类绑定的方法重名，当子类调用该重名方法时，会先调用继承父类的方法，然后再调用自己绑定的方法
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

// 引用结构体的方法，引用传递，会改变原有结构体的值
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

// 值结构体的方法，值传递，不会改变原有结构体的值
func (a Russian) Speak() {
	fmt.Println("俄罗斯人" + "说" + "要打美国人")
}
func (a Russian) Eat() {
	fmt.Println("俄罗斯人" + "吃" + "什么？")
}

// 当实现了Human接口的结构体就可以调用Tell函数了
func Tell(ren Human) {
	fmt.Println("好的", ren)
}
func Tellhao(who interface{}) {
	h := who.(Human) // 断言，传进的任何类型只有是Human类型的才能实现Tell()函数
	Tell(h)
}

func main() {
	// 这里就看出了多态的，不同的对象类型，调用相同方法，执行的是不同函数
	a := &American{
		"jieli",
		"English",
		"汉堡",
	}
	r := &Russian{
		"达瓦里氏",
		"Russian",
		"红肠",
	}
	a.Speak()
	a.Eat()
	r.Speak()
	r.Eat()

	// 当类都实现了该接口中的方法时，该类的对象是该自定义的接口类型
	var c Human
	c = r
	c.Speak()
	c.Eat()
	Tellhao(c)
}
