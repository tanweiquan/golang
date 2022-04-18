package main

import (
	"fmt"
	"reflect"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力
/*
变量包含类型信息和值信息
类型信息：是静态的元信息，是预先定义好的 (获取类型信息：reflect.TypeOf，是静态的,返回一个装载着其动态类型的reflect.Type)
值信息：是程序运行过程中动态改变的 (获取值信息：reflect.ValueOf，是动态的，返回一个装载着其动态值的reflect.Value)
*/

// 空接口与反射
// 反射获取interface类型信息
func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("TypeOf返回值：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println("Kind返回值：", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是float64类型")
	case reflect.String:
		fmt.Println("a是string类型")
	}
}

// 反射获取interface值信息
func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("ValueOf返回值：", v)
	k := v.Kind()
	fmt.Println("Kind返回值：", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

// 反射修改interface值信息
func reflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("ValueOf返回值：", v)
	k := v.Kind()
	fmt.Println("Kind返回值：", k)
	if k != reflect.Ptr {
		fmt.Println("不是指针类型,没法进行修改操作")
		return
	}
	// Elem()获取地址指向的值
	v.Elem().SetFloat(7.9)
	fmt.Println("case:", v.Elem().Float())
	// 地址
	fmt.Println(v.Pointer())
}

// ..............................................
// 结构体与反射
// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello(p1, p2 string) {
	fmt.Println("Hello", p1, p2)
}

// 查看类型、字段和方法
func poni(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型:", t)
	fmt.Println("字符串Name类型:", t.Name())
	v := reflect.ValueOf(a)
	fmt.Println("值:", v)
	// 获取属性(结构体字段个数，获取字段对应的值,结构体绑定的方法个数)
	h := t.Field(0)               // 取对应索引的字段（字段名、字段类型）
	nn := t.NumField()            // 结构体字段个数
	val := v.Field(0).Interface() // 获取索引字段对应的值
	m := t.NumMethod()            // 结构体绑定的方法个数
	r := t.Method(0)              // 结构体绑定的方法（方法名以及方法类型）
	fmt.Println(h.Name, h.Type)
	fmt.Println(nn)
	fmt.Println(val)
	fmt.Println(m)
	fmt.Println(r.Name, r.Type)
}

// 修改结构体的值
// 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 根据字段名取字段值
	f := v.FieldByName("Name")
	fmt.Println(f)
	if f.Kind() == reflect.String {
		// 修改字段对应的值
		f.SetString("kuteng")
	}
}

// 调用方法
func callmethod(u interface{}) {
	v := reflect.ValueOf(u)
	// 根据方法名获取方法
	m := v.MethodByName("Hello")
	fmt.Printf("%T\n", m)
	// 构建一些参数
	param1 := "这是自定义字符串参数1"
	param2 := "这是自定义字符串参数2"
	rf1 := reflect.ValueOf(param1)
	rf2 := reflect.ValueOf(param2)
	args := []reflect.Value{rf1, rf2} // 没参数的情况下：var args []reflect.Value
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

// 获取字段的tag
type Student struct {
	Name string `json:"name1" db:"name2"` // 保证json和db的字段按规定的字段转义
}

func gettag(s interface{}) {
	v := reflect.ValueOf(s)
	// 获取对应字段类型信息
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	// 根据标签获取转义后的字段
	t1 := f.Tag.Get("json")
	t2 := f.Tag.Get("db")
	fmt.Println(t1)
	fmt.Println(t2)
}
func main() {
	var x float64 = 3.4
	reflectType(x)
	reflectValue(x)
	// 反射修改值要传指针
	reflectSetValue(&x)
	// 结构体
	y := User{16, "小华", 26}
	poni(y)
	SetValue(&y)
	fmt.Println(y)
	callmethod(y)
	var s Student
	gettag(&s)

}
