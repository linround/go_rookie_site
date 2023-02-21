package main

import "fmt"

type Test struct {
	name string
}

func main() {
	//f1()
	f2()
}
func f1() {
	fmt.Println("修改字符串的第一个字符")
	//如何修改字符串中的第一个字符
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	//
	fmt.Println(s2)
	fmt.Println("获取字符串的子串")
	subStr := str[2:4]
	fmt.Println(subStr)
	fmt.Println("使用for循环遍历一个字符串")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%q", str[i])
	}
	println()
	fmt.Println("使用for-range循环遍历一个字符串")
	// 索引、值
	for _, ch := range str {
		fmt.Printf("%q", ch)
	}
}

func f2() {
	var arr1 = []int{10, 20, 30, 40, 50, 60}
	for _, value := range arr1 {
		fmt.Printf("%d ", value)
	}
	println(" \n")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d ", arr1[i])
	}
	println("\n ")
	var found int
	for _, value := range arr1 {
		if value == 20 {
			found = value
		}
	}
	println(" \n")
	fmt.Println("发现的值：", found)
}
