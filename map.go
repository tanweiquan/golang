package go_map

import "fmt"

// 引用类型：特指slice、map、channel这三种预定义类型

// go语言中提供的映射关系的容器为map，其内部使用散列表(hash)实现的
// map是一种无序的基于key-value的数据结构，go语言中的map是引用类型，必须初始化才能使用
// 语法：map[KeyType]ValueType，其中KeyType表示键的类型，ValueType表示键对应的值的类型
func main() {
	var a map[string]int
	fmt.Println(a == nil)        // true,还没初始化，还没在内存开辟空间
	a = make(map[string]int, 10) // 要估算好容量，避免在程序运行时再动态扩容
	a["理想"] = 18                 // 赋值
	a["王子"] = 25                 // 赋值
	fmt.Println(a)               // map[王子:25 理想:18]
	fmt.Println(a["理想"])         // 18
	// 用OK接收返回的布尔值
	value, ok := a["娜扎"] // 如果不存在这个key,value接收值的对应的类型的零值，ok接收返回布尔值的零值false
	if !ok {             // !ok对应是true，布尔表达式是true的，符合执行条件，执行下面表达式
		fmt.Println("查无此key")
	} else {
		fmt.Println(value) // 查无此key
	}
	// map的遍历(for range遍历)
	for k, v := range a {
		fmt.Println(k, v) // 王子 25 ; 理想 18
	}
	// 只遍历key
	for k := range a { // 注意当只遍历前面的值时，后面的值不需要用匿名变量接收
		fmt.Println(k) // 王子，理想
	}
	// 只遍历value
	for _, v := range a { // 用匿名变量接收
		fmt.Println(v) // 25,18
	}
	// 删除
	delete(a, "王子")
	fmt.Println(a)  // map[理想:18]
	delete(a, "沙河") // 删除不存在的key
	fmt.Println(a)  // 不影响，打印结果仍为map[理想:18]
	// map和slice(切片)的组合
	// 元素类型为map的切片
	var s = make([]map[string]int, 5) // 这里是切片初始化了，内部的map还没初始化
	s[0] = make(map[string]int, 1)    // 这里初始化了s的第一个元素为map类型
	s[0]["沙河"] = 10                   // [map[沙河:10] map[] map[] map[] map[]]
	fmt.Println(s)
	// 值为切片类型的map
	var m = make(map[string][]int, 10) // 这里是map初始化了，切片还没初始化
	m["北京"] = []int{10, 20, 30}        // 切片初始化
	fmt.Println(m)                     // map[北京:[10 20 30]]
}
