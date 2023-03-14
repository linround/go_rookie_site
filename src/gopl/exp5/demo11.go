package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("接收到的错误信息：", p)
		}
	}()
	fmt.Println("main正常执行：，main")
	panic("mainPanic")
}
