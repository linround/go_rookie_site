package main

import (
	"fmt"
	"time"
)

// 信号量模式

func compute(ch chan int) {
	ch <- someComputation()
}
func someComputation() int {
	time.Sleep(5e9)
	return 100
}

func main() {
	println("信号量模式开始")
	ch := make(chan int)
	println("进入等待")
	go compute(ch)
	println("等待结束")
	doSomethingElseForWhile()
	println("接收传递的信息")

	// 运行到此处的时候会一直等待信号数据
	result := <-ch
	println("收到了对应数据")
	fmt.Println(result)
}
func doSomethingElseForWhile() {

}
