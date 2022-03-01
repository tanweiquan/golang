package main

import "fmt"

// 数组array:存放元素的容器
// 数组必须要指定容器里元素的类型和长度。（另外当长度不确定时，可用[...]让程序自行推断长度。

// 元素初始化
var x = [3]bool{false, true, true}
var y [3]bool // 如果不初始化，使用元素类型的默认零值
var z = [...]int{1, 2, 3, 4, 5}

// 多维数组
var n = [2][3]int{
	{1, 2, 3},
	{6, 7, 8},
}

func main() {
	a := [5]int{1, 3, 5, 7, 9}
	b := [4]string{"hello", "沙河", "hi", "沙僧"}

	// 用索引指定不同位置的值
	m := [6]int{0: 66, 1: 88} // 指定位置的下标为0的值为66，下标为1的值为88，其他位置的值使用默认零值0
	fmt.Printf("%v\n %v\n %v\n %v\n %v\n %v\n %v\n", a, b, x, y, z, m, n)
	fmt.Println(a[0]) // 打印索引为0的值

	// 数组是值类型，根据索引改变副本不会改变源数组
	// 改变源数组只能照源数组索引改变值
	copy := a
	copy[0] = 100
	fmt.Println(a[0]) // a[0]没改变
	a[0] = 888
	fmt.Println(a[0]) // a[0]改变了

	// 遍历
	// for 遍历，根据索引遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// for i,v := range 数组{} 来遍历所有的元素
	for i, v := range a {
		fmt.Println(i, v)
	}
	// 多维数组的遍历,使用for i,v := range 数组{}遍历
	for i1, v1 := range n {
		fmt.Printf("第一层数组的索引为%v---值%v\n", i1, v1)
		for i2, v2 := range v1 {
			fmt.Printf("第二层数组的索引为%v---值%v\n", i2, v2)
		}
	}

}
