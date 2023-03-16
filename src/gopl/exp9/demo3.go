package main

import "fmt"

var sema = make(chan string, 1)

func main() {
	sema <- "结束"
	fmt.Println("开始")
	fmt.Println(<-sema)
}
