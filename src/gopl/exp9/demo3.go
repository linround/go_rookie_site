package main

import (
	"fmt"
	"time"
)

var sema = make(chan string, 0)

func routine() {
	for {
		fmt.Println("routine 开始")
		// 在这里等待 信号时会处于阻塞
		// 接收到信号后再次执行，输出信号后继续处于阻塞
		fmt.Println(<-sema)
		fmt.Println("routine 结束")
	}
}
func main() {
	go routine() // 这个routine程序会一直等待（如果没有缓存区，一点要在发送之前准备好goroutine）
	sema <- "结束"
	// 如果没有缓冲去在这里就会阻塞，
	// 没有 d对于没有缓存的通道 可以看作是一条单纯的连接通道
	fmt.Println("开始")
	time.Sleep(1 * time.Second)
	//fmt.Println(<-sema)
}
