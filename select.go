package main

import (
	"fmt"
	"time"
)

// select多路复用
// 使用场景:同时从多个通道接收数据。
// 问题：通道在接收数据时，如果没有数据可以接收将会发生阻塞
/*
1、select可以同时监听一个或多个channel，直到其中一个channel ready
2、如果多个channel同时ready，则随机选择一个执行
3、可以用于判断管道是否存满
select {
case <-chan1:
   // 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
   // 如果成功向chan2写入数据，则进行该case处理语句
default:
   // 如果上面都没有成功，则进入default处理流程
}
*/
func main() {
	// select监听channel
	t := make(chan int, 1)
	u := make(chan int, 1)
	go func() {
		t <- 16
	}()
	go func() {
		u <- 17
	}()
	time.Sleep(time.Second)

	select {
	case x := <-t:
		fmt.Println(x)
	case y := <-u:
		fmt.Println(y)
	default:
		fmt.Println("没有值")
	}
	close(t)
	close(u)
	// select 判断通道是否已经满了
	m := make(chan string, 1)
	defer close(m)
	time.Sleep(time.Second)
forend:
	for {
		select {
		case m <- "a":
			fmt.Println("通道没满")
		default:
			fmt.Println("通道已经满了")
			break forend
		}
	}
	fmt.Println("运行结束")

}
