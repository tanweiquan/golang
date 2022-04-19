package go_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// goroutine 只是由官方实现的超级"线程池"

/* Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。
Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine。
当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。
Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。 */

/* 在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束
创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行的
主协程退出，子协程也跟着退出了。 golang main函数结束，所有协程都被结束了 */
func main() {
	go func() {
		fmt.Println("hello,Sleep")
	}()
	// 防止退出
	// 1.Sleep
	time.Sleep(time.Second)
	// 2.等待组
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("hello,WaitGroup")
		wg.Done()
	}()
	wg.Wait()
	/*
		1.一个操作系统线程对应用户态多个goroutine。
		2.go程序可以同时使用多个操作系统线程。
		3.goroutine和OS线程是多对多的关系，即m:n。
	*/
	// runtime包
	// runtime.Gosched()：马上让出当前正在运行的CPU时间片，重新等待安排任务
	go func() {
		fmt.Println(1)
		runtime.Gosched()
	}()
	go func() {
		fmt.Println(2)
	}()
	time.Sleep(time.Second)
	// runtime.Goexit()：马上结束并退出当前正在运行的协程
	go func() {
		fmt.Println("A")
		defer fmt.Println("B")
		runtime.Goexit()
		fmt.Println("C") // 会打印A、B,不会打印C
	}()
	time.Sleep(time.Second)
	// runtime.GOMAXPROCS：确定需要使用多少个OS线程来同时执行Go代码（默认值是机器上的CPU核心数）
	// runtime.GOMAXPROCS(1) // 将逻辑核心数设为1，此时是做完一个任务再做另一个任务
	runtime.GOMAXPROCS(2) // 将逻辑核心数设为2，此时两个任务并行执行
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("X", i)
		}
	}()
	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println("Y", j)
		}
	}()
	time.Sleep(time.Second)

}
