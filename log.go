package main

import (
	"fmt"
	"log"
	"os"
)

// 标准库log

func main() {
	v := `这是一条日志`

	// 使用Logger，默认情况下的logger只会提供日志的时间信息
	// 普通打印
	log.Print(v)
	log.Println(v)
	log.Printf("%s\n", v)
	// 打印触发错误的信息,打印完成后会调用os.Exit(1),退出程序
	// log.Fatal(v)
	// log.Fatalln(v)
	// log.Fatalf("%s\n", v)
	// 打印触发崩溃的信息，打印完成后触发panic退出程序
	// log.Panic(v)
	// log.Panicln(v)
	// log.Panicf("%s\n", v)

	// 配置logger，格式化输出更多信息

	// 配置Flag选项·
	// log库中的Flag选项有许多常量：
	// Ldate:日期
	// Ltime:时间（秒）
	// Lmicroseconds:时间（微妙）
	// Lshortfile:文件名+行号
	// Llongfile:文件全路径名+行号
	// LUTC:使用UTC时间
	// LstdFlags:标准logger的初始值
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ltime) // 设置logger的输出选项
	log.Println(v)

	// 配置日志信息前缀 （两个方法）
	// Prefix函数用来查看标准logger的输出前缀,即Flag选项中的常量(包含不设置时的默认选项)
	// SetPrefix函数用来设置输出前缀
	s := log.Prefix()
	log.Print(s)
	log.SetPrefix("这是log前缀")
	log.Println(v)

	// 配置日志输出位置：SetOutput()
	// 1.新建并打开文件
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	// 2.设置输出到文件中
	log.SetOutput(logFile)
	// 3.设置日志的选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// 4.设置日志的前缀
	log.SetPrefix("[前缀]")
	// 如果要每次输出日志就配置相关日志信息，可以写入init函数里
	// init()-->log.Print
	log.Print("这是一条日志")

	// 创建logger：New()
	logFilex, err := os.OpenFile("./xy.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logger := log.New(logFilex, "<前缀>", log.Ltime|log.Llongfile)
	// 如果要每次输出日志就配置相关日志信息，可以写入init函数里
	// init()-->log.Print
	logger.Print("这是一条日志")

	// Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，
	// 我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等
}
