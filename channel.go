package go_channel

import (
	"fmt"
	"time"
)
// 总结：golang的chan翻译成中文就是管道，顾名思义，就是管道的一端用来读，另一端用来写，这与write和read函数的性质是非常相似的，比如说管道中没数据，就会发生读阻塞，管道中数据是满的，就会发生写阻塞。又类似生产者和消费者，也就是必须有生产（写），才能消费（读）
// chan是用来进行go程间通信的，chan的相关操作有发送、接收、关闭，管道是一种引用类型
// len函数可以获取通道内元素的数量，cap函数获取通道的容量
// 注意：
// 写和读一般放在2个go程，否则容易死锁
// 如果阻塞发生在主go程，并且没有其他goroutine在运行，会导致程序挂掉
// 当管道中的数据读完后，再继续读，存在以下两种情况：
// 1、通道已经close，则读出的是通道内类型的零值
// 2、通道未close，会发生阻塞

// 单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
// 共享内存在不同的goroutine中容易发生竞态问题,为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
// Go语言的并发模型是CSP，提倡通过通信共享内存而不是通过共享内存而实现通信。
// Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

// 基本使用
func chanabout() { // 这里只是举例，正常情况同一个chan的发送和接收不在一个goroutine
	// 创建channel
	var ch chan int // 未初始化，ch为nil值，并不是零值
	ch = make(chan int, 1)
	// 发送
	ch <- 10
	// 接收
	x := <-ch // 另外还有<-ch接收并扔掉
	fmt.Println(x)
	// 关闭
	close(ch)
}

// 无缓冲通道
// A、如果发送操作先执行，发送方的goroutine将阻塞，直到另一个goroutine在该通道上执行接收操作
// B、如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上执行发送操作
// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化
func f1(ch chan int) {
	// 发送
	ch <- 12
}
func f2(ch chan int) {
	// 接收
	x := <-ch
	fmt.Println(x)
}

/*
将通道作为函数参数时，可以使用单向通道
1.chan<- T是一个只能发送的通道(只写通道)，可以发送但是不能接收；
2.<-chan T是一个只能接收的通道(只读通道)，可以接收但是不能发送。
*/
func chanx(ch chan<- int) {
	var x = 16
	ch <- x
}
func chany(ch <-chan int) {
	y := <-ch
	fmt.Println(y)
}

func main() {
	chanabout()
	var chx chan int
	chx = make(chan int)
	go f1(chx)
	f2(chx)
	time.Sleep(time.Second)
	close(chx)
	fmt.Println(".......................")

	// 如何优雅关闭chan
	// 当关闭chan后,发送方不能发送值，否则引发panic，接受方还可以继续接收值,直到接收该类型的零值为止
	// 那如何判断chan已经关闭呢？
	/*
		1、可以利用ok=false
		A、读取关闭后的无缓存通道，不管通道中是否有数据，返回值都为0和false。
		B、读取关闭后的有缓存通道，将缓存数据读取完后，再读取返回值为0和false。
	*/
	chy := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			chy <- i
		}
		close(chy)
	}()
	go func() {
		for {
			y, ok := <-chy
			fmt.Println("y:", y) // 打印0~9
			if !ok {
				fmt.Println("通道已经关闭")
				fmt.Println(y) // 打印该类型的零值后就退出
				break
			}
		}
	}()
	time.Sleep(time.Second)
	/*
		2、for value := range ch {}
		往通道发送完数据后，必须关闭通道，否则range遍历会出现死锁。
		通道关闭后会退出for range循环
	*/
	chz := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			chz <- i
		}
		close(chz)
	}()
	for i := range chz {
		fmt.Println("i:", i) // 打印0~9
	}
	time.Sleep(time.Second)
	// 单向通道
	cht := make(chan int, 1)
	go chanx(cht)
	chany(cht)
	time.Sleep(time.Second)
}
