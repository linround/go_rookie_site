package main

import "fmt"

func main2() {
	strings := []string{"google", "run"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	// 这里定义了6个空间的数组
	// 未填充的默认值为0
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Println(i, x)
	}
}

func main1() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	// go关键字有点类似async的意思
	main2()
	go main1()
	fmt.Println(sum)
}
