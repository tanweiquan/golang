package new_make

import "fmt"

func main() {
	// 关键字 new 主要用来给基本类型和结构体开辟内存，返回对应类型的指针。
	// 关键字 make 是用来给特殊类型（slice、map、chan）开辟内存，返回对应的类型。初始化后我们可以变量赋值。
	// new
	a := "hello world"
	b := new(string)
	b = &a
	fmt.Println(b)
	// make
	// 有值的初始化
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Println(e, f)
	// 无值初始化，使用make来创建和初始化切片slice和字典map
	s := make([]int64, 5)
	s1 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	fmt.Printf("%#v,len:%#v\n", m2, len(m2))
}
