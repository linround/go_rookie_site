package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 以下是两个goroutine在执行
	// goroutine 是轻量级线程，goroutine 的调度是由Golang 运行时进行管理
	// 可以使用go语句开启一个新的运行期线程
	// 同一个程序中所有goroutine共享一个地址空间
	go say("world")
	say("hello")
}
