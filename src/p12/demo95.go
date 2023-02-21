package main

import (
	"fmt"
	"time"
)

// 通道工厂模式
// 不讲通道作为参数传递，使用函数来生成一个通道并返回
// 函数内有个匿名函数被协程调用

func main() {
	stream := pump2()
	go suck2(stream)
	// 使用sleep等待结束
	time.Sleep(1e9)
}
func pump2() (ch chan int) {
	ch = make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return
}
func suck2(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
