package go_if_switch

import "fmt"

// 条件语句---if
// 基本格式：
/*
  if 变量声明；判断表达式1的布尔值是true{
  	分支1
  }else if 判断表达式2的布尔值是true{
  	分支2
  }
  else {
  	分支3
  }
*/

// 条件语句---switch
// 基本格式：
/* switch 变量声明；变量{
   case  判断表达式1为true：
    分支1
   case  判断表达式2为true：
    分支2
   case  判断表达式3为true：
    分支3
   ...
   default：
    分支n
}
*/

func main() {
	// if例子：
	age := 19
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年人")
	} else {
		fmt.Println("快写作业")
	}
	// 作用域:只在条件语句中生效
	if a := 19; a > 18 {
		fmt.Println("a>18")
	} else {
		fmt.Println("a<18")
	}
	// switch例子：
	// 作用域:只在条件语句中生效
	switch m := 3; m {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: //默认值
		fmt.Println("无效的数字")
	}
	// 设置多个元素在case集合中
	// 只要变量符合该case集合中任意一元素，则为true
	switch x := 5; x {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default: //默认值
		fmt.Println(x)
	}
	// 变种1：
	switch age := 30; {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default: //默认值
		fmt.Println("活着真好")
	}

	// 变种2：
	n := 4
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default: //默认值
		fmt.Println("无效的数字")
	}
	// 变种3：
	yage := 61
	switch {
	case yage < 25:
		fmt.Println("好好学习吧")
	case yage > 25 && yage < 35:
		fmt.Println("好好工作吧")
	case yage > 60:
		fmt.Println("好好享受吧")
	default: //默认值
		fmt.Println("活着真好")
	}

}
