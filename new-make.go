package new_make

import "fmt"

func main() {
	// 关键字 new 主要用来创建一个引用类型的结构体，只有结构体可以用。
	// 关键字 make 是用来创建和初始化一个slice或者map。我们可以直接赋值来使用
	// 直接赋值：
	e := []int64{1, 2, 3}                 // slice
	f := map[string]int64{"a": 3, "b": 4} // map
	fmt.Println(e, f)
	// 使用make来创建和初始化切片slice和字典map
	s := make([]int64, 5)
	s1 := make([]int64, 0, 5)
	m1 := make(map[string]int64, 5)
	m2 := make(map[string]int64)
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	fmt.Printf("%#v,len:%#v\n", m2, len(m2))
}
