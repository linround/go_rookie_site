package main

import (
	"fmt"
	"os"
	"runtime"
)

// 省略类型说明符 编译器根据变量的值来推断其类型

var r = 100

const pi = 3.17

func main() {
	const (
		a = iota
		b = iota + 63
		c
	)

	var (
		home = os.Getenv("HOME")
		goos = runtime.GOOS
		path = os.Getenv("PATH")
	)
	home = "145"
	fmt.Println(a, b, c, home, goos)
	fmt.Println(path)
}
