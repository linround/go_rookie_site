package main

import "fmt"

func main() {
	var s []int // 定义切片
	s[0] = 100
	s = append(s, 100)
	x := s // 值传递
	x[0] = 10000
	fmt.Print(s)
	fmt.Print(x)
	println()
	sl()
}

func sl() {
	// 切片自身就是指针
	var arr1 [6]int // 定义一个容量为6的切片
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Print(arr1)
	println()
	fmt.Print(slice1)
	println()
	fmt.Print(cap(slice1))
	slice1 = slice1[0:4]
	println()
	fmt.Print(slice1)
	// 以下可以证明slice1就是一个指向数组arr1的指针
	slice1[0] = 45

	fmt.Print(arr1)

}
