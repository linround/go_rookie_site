package main

import (
	"fmt"
	"time"
)

// go协程意味着并行（或则可以以并行的方式部署）
// go协程通过通道来通信；协程通过让出和恢复操作来通信
func main() {
	fmt.Println("main函数执行")
	// 如果移除了go关键字
	// 就不会产生协程
	go longWait()
	go shortWait()
	fmt.Println(" main 开始等待")
	// sleep 可以按照指定时间来暂停函数或协程的执行
	// 如果我们不在main中等待
	// 整个程序就会暂停执行，协程也会随之消亡
	// main 函数退出的时候 他不会等到任何其他非main协程的结束
	time.Sleep(10 * 1e9) // 10秒时间
	fmt.Println("main 结束")

}
func longWait() {
	fmt.Println("开始longWait")
	time.Sleep(5 * 1e9) // 5秒时间
	fmt.Println("longWait结束")
}

func shortWait() {
	fmt.Println("开始shortWait")
	time.Sleep(2 * 1e9) // 2秒时间
	fmt.Println("shortWait结束")
}
