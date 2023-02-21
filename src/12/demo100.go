package main

import (
	"fmt"
	"time"
)

var num int = 0

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump7(ch1)
	go pump8(ch2)
	go suck8(ch1, ch2)
	time.Sleep(1e9)
	fmt.Println("最终次数：", num)
}
func pump7(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 100
	}
}
func pump8(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 5
	}
}
func suck8(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			num++
			fmt.Println("通道一的值", v)
		case v := <-ch2:
			num++
			fmt.Println("通道二的值", v)
		}

	}
}
