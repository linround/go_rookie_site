package main

import (
	"fmt"
	"time"
)

// 闭包和协程的使用

var values = []int{1, 2, 3, 4, 5, 6}

func a() {
	//版本A
	for ix := range values {
		func() {
			fmt.Print(ix, "  ")
		}() // 使用闭包打印每个索引值
	}
	println("版本A")
}
func b() {

	//	版本B
	// 将闭包当作协程进行调用
	// 这里的ix实际是一个单变量
	// 表示每个数组元素的索引值
	// 因为这些闭包只绑定到一个变量
	// 当运行这段代码的时候
	// 协程可能在循坏结束后还没有开始执行
	for ix := range values {
		go func() {
			fmt.Print(ix, "  ")
		}()
	}
	time.Sleep(4e9)
	println("版本B")
}
func c() {
	// 版本C
	// 正确的处理方式
	for ix := range values {
		go func(ix int) {
			fmt.Print(ix, "  ")
		}(ix)
	}
	println("版本C")
	time.Sleep(5e9)
}
func d() {

	//	版本D 输出值
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, "  ")
		}()

	}
	time.Sleep(5e9)
	println("版本D")
}
func main() {
	//a()
	b()
	//c()
	//d()
}
