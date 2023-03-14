package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println("堆栈信息输出开始")
	os.Stdout.Write(buf[:n])
	fmt.Println("堆栈信息输出结束")
}
