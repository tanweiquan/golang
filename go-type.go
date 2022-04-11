package go_type

import (
	"fmt"
	"strings"
)

// 当变量没初始化，则变量使用类型的默认零值
// 值类型：int、float、bool和string这些类型都属于值类型
// 所有赋值过程都是值拷贝
// 整型：
// 有符号int (有正负，范围：-(2^(n-1))至2^(n-1)-1)
// 无符号uint(只有正，范围：0~2^n)
// 32位系统用int32或uint32，64位系统用int64或uint64

// 浮点型：支持float32和float64
// float32：浮点数的最大范围为3.4e38
// float64：浮点数的最大范围为1.8e308

// 整型变量和浮点型变量的默认零值为0
var a int
var b int8 = -4
var c uint64 = 12

var e float32
var f float64 = 16.1

// 布尔类型
// 布尔型变量默认零值为false
var m bool
var x, y = false, true

// string类型
// string类型变量的默认零值为空字符串
var s string
var s1 = "hello_world"

/*

\本来就有特殊含义的，我应该告诉程序我写的\就是单纯的\
	\r 回车符
    \n 换行符
    \t 制表符（水平制表）
    \' 使用单引号'
    \" 使用双引号"
    \\ 反斜杠\
*/
var path = "\"F:\\WeChat\\locales\""

// 字符，单引号'包裹
// 字符编码历史：ASCII->unicode->UTF-8,GO使用的是UTF-8编码的
// 一个英文字符在UTF-8只占用一个字节(与英文字符在ASCII上一致)，UTF8编码下一个中文汉字由3~4个字节组成
// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符
// rune类型，代表一个 UTF-8字符
// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组
// 下面这种声明方式只会将字符的码值赋值给变量,相当于var s2 = byte('w')
var s2 = 'w'

func main() {
	fmt.Println(a, b, c)
	fmt.Println(m, x, y)
	fmt.Println(e)
	// 简短声明32位浮点型
	g := float32(16.1)
	fmt.Printf("类型：%T,值：%v\n", g, g) // 类型：float32,值：16.1
	// 简短声明不指定类型会使用系统的位数的浮点型
	h := 16.1
	fmt.Printf("类型：%T,值：%v\n", h, h) // 类型：float64,值：16.1
	// 注意不同类型的变量是不能比较的，也就是说g和h是不能比较大小的
	fmt.Println(s)    //
	fmt.Println(s1)   // hello_world
	fmt.Println(path) // "F:\WeChat\locales"
	fmt.Println(s2)   // 输出码值119

	f1 := "hello,沙河。"
	// len()计算字符串字节数
	fmt.Println(len(f1)) // 15
	// 相当于以下这种，计算字符串所占字节数
	sb := []byte(f1)
	fmt.Println(len(sb)) // 15 --->5+1+2*3+3
	// 计算字符串里的字符数量时，可用rune()包裹起来
	sp := []rune(f1)
	fmt.Println(len(sp)) // 9

	// 拼接字符串
	f2 := "hello,小溪。"
	f3 := f1 + f2
	fmt.Println(f3) // hello,沙河。hello,小溪。
	// 字符串切割
	f4 := strings.Split(f2, ",")
	fmt.Println(f4) // [hello 小溪。]
	// 字符串切片中的字符拼接
	f5 := strings.Join(f4, "+")
	fmt.Println(f5) // hello+小溪。
	// 字符串修改
	// 先转换位rune切片
	f6 := []rune(f2)
	f6[6] = '大'
	// 再将rune切片转换回string类型
	f7 := string(f6)
	fmt.Println(f7)

}
