package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int, 0)
	out <- 2
	//go func() { out <- 2 }()
	go f1(out)
	// 是为了等待goroutine执行完成
	time.Sleep(2e9)
}
