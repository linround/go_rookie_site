package main

import (
	"fmt"
)

func generate(ch chan int) {
	// 向通道中进行传递值
	for i := 2; ; i++ {
		ch <- i
	}
}

// 协程filter
func filter(in, out chan int, prime int) {
	for {
		// 获取输入通道的值
		i := <-in
		// 取余不等于0
		if i%prime != 0 {
			// 将无法被prime整除的数字传入到输出通道
			out <- i
		}
	}
}
func main() {
	ch := make(chan int)
	go generate(ch)
	fmt.Println("开始")
	for {
		// 输入通道的值
		// 输入2
		// 分析 一开始是2
		prime := <-ch
		fmt.Println("prime:", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
