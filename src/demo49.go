package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	// 切片可以反复扩展
	// 直到占满容量
	slice1 = slice1[0 : len(slice1)+10]
	fmt.Println(cap(slice1))

}
