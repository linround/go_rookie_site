package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func main() {
	fmt.Println("开始调用test")
	test()
	fmt.Println("结束调用test")
}
func badCall() {
	panic("bad end")
}
func test() {
	// 这里的defer 调用必须在badCall前面
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("defer 错误", e)
		}
	}()

	badCall()
	fmt.Println("bad call 调用结束")
}
