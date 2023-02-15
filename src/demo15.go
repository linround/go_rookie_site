package main

import "fmt"

const MAX int = 3

// 指向整型的指针数组
func change(arr [MAX]*int) {
	// 获取第一个指针指向的地址
	// 赋值该地址值为1000
	*arr[0] = 1000
}

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}
	change(ptr)
	println(a[0])
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
