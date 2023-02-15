package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)
	fmt.Println(numbers[1:2])
	fmt.Println(numbers[:4])
	fmt.Println(numbers[0:5])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	numbers1 = append(numbers1, 1)
	fmt.Println(numbers1)
	fmt.Println(cap(numbers1))
	numbers1 = make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers1)
	fmt.Println(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
