package main

import "fmt"

// 以下程序取得10以下数字的平方
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		// 只需要告诉接收者goroutine,所有数据已发送完毕
		// 无论一个channel是否被关闭
		// 当他没有被引用时将会被回收
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
