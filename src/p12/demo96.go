package main

import (
	"fmt"
	"time"
)

// 给通道使用for循环
func main() {
	suck3(pump3())
	time.Sleep(1e9)
}
func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 通道迭代
func suck3(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println("suck3:", v)
		}
	}()
}
