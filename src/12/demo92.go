package main

import (
	"fmt"
)

func f1(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func fp(ch chan int) {
	for i := 0; ; i++ {

		ch <- i
	}
}

func main() {
	out := make(chan int, 100)
	go fp(out)
	go f1(out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	// 放在这里也不会执行，因为通道关闭了
	//go f1(out)
}
