package main

import (
	"fmt"
	"time"
)

// 本质上是生产者消费者模型
// 可以有效控制goroutine数量，防止暴涨

// 例子:计算10个数的相加结果

func createnum(i int, m chan<- int) {
	m <- i
}
func countnum(m <-chan int) {
	y := <-m
	fmt.Println(y)
}
func main() {
	// 非goroutine池，有多少个任务进来就开启多少个goroutine
	x := make(chan int, 1)
	var y int
	for i := 0; i < 10; i++ {
		go createnum(i, x)
		y++
	}
	j := 0
	for {
		go countnum(x)
		if j > y {
			break
		} else {
			j++
			continue
		}
	}
	// goroutine池，可以自定义goroutine数量，当某个goroutine先处理完接到的任务后还可以再接任务
	t := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go createnum(i, x)
	}
	for i := 0; i < 10; i++ {
		go countnum(x)
	}
	time.Sleep(time.Second)
}
