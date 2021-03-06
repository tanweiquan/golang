package type_change

import (
	"fmt"
	"strconv"
)

// 不同类型的变量不能相互赋值
// var i int = 9;type myint int;var a myint;这里是自定义类型，i不能赋值给a
// var i int = 9;type myint = int;var a myint;a=i;这里是类型别名，可以相互赋值
func main() {

	s := "hello,小溪"
	f1 := []rune(s)
	f2 := []byte(s)
	// string()只支持将[]rune()和[]byte()转换为string类型
	fa1 := string(f1)
	fa2 := string(f2)
	fmt.Println(fa1, fa2)
	// 基本数据类型之间的转换
	t1 := 10
	t2 := 10.1
	t3 := true
	// 整型与浮点型之间的转换
	t4 := float64(t1)
	fmt.Println(t4) // 10
	// int转int64
	t5 := int64(t1)
	// 基本数据类型转string型
	// 法1：使用fmt.Sprintf()
	fo := fmt.Sprintf("%v %v %v\n", t1, t2, t3)
	fmt.Println(fo) // 10 10.1 true

	// 法2：使用strconv包的函数
	// int转string类型
	ft1 := strconv.Itoa(t1)
	// int64转string类型
	ft2 := strconv.FormatInt(t5, 10) // 10为进制
	// float64转string类型
	ft3 := strconv.FormatFloat(t2, 'f', 10, 64) //'f'是格式 10表示小数保留十位 64表示这个小数是float64
	// bool类型转string类型
	ft4 := strconv.FormatBool(t3)
	fmt.Println(ft1, ft2, ft3, ft4) // 10 10.1000000000 true

	fu1 := "16"
	fu2 := "16.1"
	fu3 := "true"
	// string类型转基本数据类型
	fs1, _ := strconv.Atoi(fu1)
	fs2, _ := strconv.ParseFloat(fu2, 64) // 64表示这个小数是float64
	fs3, _ := strconv.ParseBool(fu3)
	fmt.Println(fs1, fs2, fs3)
	// 类型推断
	i := 6
	x := interface{}(i) // 使用x.(type)前一定要转换为interface{}类型
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("不合法")
	}
	xtype, ok := x.(int)
	fmt.Println("i的值为：", xtype)
	fmt.Println("i的类型是int,对还是错", ok)

}
