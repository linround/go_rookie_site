package main

import (
	"fmt"
)

type ii interface {
}

type ss struct {
}

func main() {
	var i ii
	if i == nil {
		fmt.Println("空接口类型")
	}
	fmt.Println(i)
	var s ss
	fmt.Println(s)
}
