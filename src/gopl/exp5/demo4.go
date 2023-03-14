package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "abcd"
	sr := strings.Map(func(r rune) rune {
		return r + 1
	}, s)
	fmt.Println(sr)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
