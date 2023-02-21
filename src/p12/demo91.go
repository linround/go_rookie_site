package main

import "fmt"

// 默认情况下 通信是同步且无缓冲的；
func main() {
	ch1 := make(chan int)
	// 开启一个协程 来发送数据
	go pump(ch1)
	// 开启另一个协程 来接收数据
	go suck(ch1)

	// 非循环 只能接收有限次的通道数据
	// 由于以下需要运行一点时间，所以 协程的运算会持续
	// 如果没有以下程序执行，协程会立即结束
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

// 定义一个函数，该函数同样也使用无限循环 来接受通道中的数据
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}

}

// 一个无限循坏使用通道进行传输数据
func pump(ch chan int) {
	for i := 0; ; i++ {

		ch <- i
	}
}
