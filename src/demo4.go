package main

import (
	"fmt"
	"unsafe"
)

// 使用const 声明常量
const a string = "abc" //显式类型定义：
const b = "abc"        // 隐式类型定义

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = WIDTH * LENGTH
	fmt.Print("面积为：", area)
	println()
	fmt.Printf("面积为：%d", area)
	println()
	println("面积为:", area)
	println()
	const (
		c = unsafe.Sizeof(a)
	)

	println(c)
}
