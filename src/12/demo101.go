package main

import (
	"fmt"
	"time"
)

// time.Ticker
// 这个对象以指定的时间间隔 重复的向通道c发送时间值

func main() {
	tick := time.Tick(1e8)
	// boom 只会传递一次数据
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
		default:
			fmt.Println("····")
			time.Sleep(1e9)
		}
	}
}
