package main

import "fmt"

func main() {
	var value = 100
	v2 := 45 // 不带声明格式
	fmt.Print(value, v2)
	var a string
	var num int
	var bo bool
	// nil 可以看作零值/空值
	// nil 不可以用来比较
	var a1 *int           // nil
	var a2 []int          // []
	var a3 map[string]int // map[]
	var a4 chan int       //nil
	fmt.Print(a, num, bo)
	fmt.Print(a1, a2, a3, a4)

	a5, a6, a7 := 50, 85, 89 // 并行给多个变量赋值  声明时必须是同一种变量
	fmt.Println(a5, a6, a7)

	var a8, a10 int
	var c string
	a8, a10, c = 8, 10, "c" // 字符串必须使用的是双引号
	fmt.Println(a8, c, a10)
}
