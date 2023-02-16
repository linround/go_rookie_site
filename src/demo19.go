package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var m int = 12
	var f float32 = 12
	r := m * int(f)
	i = factorial(4)
	fmt.Println(i)
	fmt.Println("")

	//	sting to int
	var n string = "15"
	s, err := strconv.Atoi(n)
	if err == nil {
		fmt.Println("正确", s*100)
	} else {
		fmt.Println("错误", err)
	}
	fmt.Println(r)

	//	 int to string

}
func factorial(n int) int {
	if n > 0 {
		return n * factorial(n-1)
	}
	return 1
}
