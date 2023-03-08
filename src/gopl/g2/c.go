package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
	fmt.Printf("%v", x)
}

var cwd string

func init() {
	// 通过单独声明的方式，防止重复声明cwd
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
