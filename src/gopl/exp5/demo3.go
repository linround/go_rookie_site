package main

import "fmt"

func add(n int) int {
	return n + n
}

var fAdd func(int2 int) int

func main() {

	println(add(10))
	if fAdd == nil {
		fmt.Println("函数空")
		return
	}

}
