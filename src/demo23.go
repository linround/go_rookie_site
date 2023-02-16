package main

import "fmt"

// 关于通道channel
// 使用通道来传递一个数据结构
// 通道可以用于两个goroutine 之间通过传递一个指定类型的值来同步运行和通讯
// 操作符 <- 用于指定通道的方向、发送或接收
// 如果未指定方向 则为双向通道
func main() {
	s := []int{7, 2, 5, 8, 9, 6}
	c := make(chan int)
	// 开启两个goroutine
	go sum(s[:len(s)/2], c)
	go sum(s, c)
	// 使用x,y来接收通道的值
	x := <-c
	y := <-c
	fmt.Println(x, y)
	main23()
	main24()
}

func main24() {
	// 关于通道的缓冲区
	// 发送端的数据可以放在缓冲区，等待接收端去获取数据，而不是立刻需要接收端去获取数据
	//	如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值
	//	如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区

	// 这里定义了一个可以存储整数类型的带缓冲通道
	// 缓冲大小为2
	ch := make(chan int, 2)
	// 由于带有缓冲，所以不必要立刻去读取数据
	//
	ch <- 1
	// 如果没有缓冲去，是没办法继续传递值
	ch <- 2
	//x := <-ch
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func main23() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	c := make(chan int)
	//	定义两个goroutine
	// 二分分别求的两边的值
	// 以下两个通道不确定谁会先执行结束
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func sum(s []int, c chan int) {
	sum := 0
	// 遍历数组
	for _, v := range s {
		sum += v
	}
	// 将结构通过通道c 进行传递
	c <- sum
}
