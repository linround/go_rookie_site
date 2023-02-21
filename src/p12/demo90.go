package main

import (
	"fmt"
	"time"
)

// 协程间的信道
// 协程是独立运行的，他们之间没有通信

// go 的特殊类型 通道 channel
// 通过通道发送类型化的数据 在协程间通信。从而避开了所有内存共享的缺陷
// 可以创建通道的通道
// 通道可以看作类型化消息队列
// 他是先进先出的结构，所有可以保证发送给他们的元素的顺序

// 运行是会检查所有的协程是否在等待
func main() {
	// 协程共享通道
	ch := make(chan string)
	go sendData(ch)
	//
	go getData(ch)
	// 等待了1秒让两个协程完成，如果不等待，sendData就没有机会输出

	time.Sleep(1e9)
	fmt.Println("接受结束")

}
func sendData(ch chan string) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
	ch <- "fff"
	ch <- "ggg"
}

// getData 使用了无限循环
// 它随着sendData 的发送完成和ch变空 也会结束
func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println("接收数据：", input)
	}
}
