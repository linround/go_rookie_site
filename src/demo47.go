package main

import "fmt"

func main() {
	var a = []string{"1", ",5"}
	for _, v := range a {
		fmt.Print(v)
	}
	// 这种
	//
	//var arr1 [5]int8  这里的arr1是一个值类型
	arr1 := new([5]int) // 这里的arr1是一个引用类型
	arr2 := arr1
	arr3 := *arr1
	arr2[0] = 100
	println()
	fmt.Print(arr1[0])
	println()
	fmt.Print(arr3)

	var arr6 = [...]int{1, 4, 5, 8}
	arr6[0] = 15555
	fmt.Print(arr6)
}
