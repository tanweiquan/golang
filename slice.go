package go_slice

import "fmt"

// 引用类型：特指slice、map、channel这三种预定义类型
//slice
//data不一定指向底层数组的起始位置
//[]int底层就是int类型底层数组，[]string底层就是string类型底层数组
//[]int的data指向底层int类型数组，len指向数组元素个数
//[]string的data指向字符串内容，len指向字节长度

//当slice追加的元素超过容量，会复制出一份底层数组，用复制的数组追加新元素并扩容成新的底层数组，然后slice指向新数组
/*扩容规则：
如果 oldCap(扩容前容量)*2 < cap(最小所需容量) --> newCap=cap
否则
     oldlen(扩容前元素个数) < 1024 --> newCap=oldCap*2
     oldlen(扩容前元素个数) > 1024 --> newCap=oldCap*1.25
*/
//所需内存=预估容量*元素类型大小
//(元素类型大小要看操作系统多少bit，64位下每个元素占8字节),实际申请到内存的是在8，16，32，48，64...里找（8分别*2、4、6...)

func main() {
	//声明一个变量是slice
	var a []int //此时还没分配底层数组，data=nil,len=0,cap=0
	//..................................................
	//通过make的方式定义变量是slice
	var b []int = make([]int, 2, 5) //此时开辟内存作为底层数组，data指向底层数组的起始地址address(Array)，len=2,cap=5
	//增加一个元素
	b = append(b, 1) //因为len=2，让前两个初始化为零值了，此时，追加的元素在底层数组的第三个
	//已经存储的元素是可以安全读写的,但超出这个范围就属于越界访问,会发生panic
	b[0] = 2 //底层数组从开始到第三个(上面已经追加)都是可以读写的，超出了第三个就属于越界访问

	//......................................................
	//new不会负责底层数组的分配，new的返回值就是slice结构的起始地址address(Slice)
	ps := new([]string) //data=nil,len=0,cap=0,此时ps就是一个地址。此时这个slice还没有底层数组
	//此时可以通过append给ps分配底层数组
	*ps = append(*ps, "eggo") //此时[]string的data(string)=eggo,len(byte)=4

	//打印值
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ps)
}
