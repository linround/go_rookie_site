package main

import "fmt"

func main() {
	// 缓冲区为10 的chan
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// 关于通道的缓冲区
	// 发送端的数据可以放在缓冲区，等待接收端去获取数据，而不是立刻需要接收端去获取数据
	//	如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值
	//	如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区

	// range 遍历 并从通道接收的数据
	// c会发送10个数据 之后就关闭了通道
	// 如果通道没有关闭，range函数就不会结束，从而在接收第11个数据的时候发生阻塞
	for i := range c {

		fmt.Println(i)

	}
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// 通道接收数据
		c <- x
		x, y = y, x+y
	}
	close(c)
}
