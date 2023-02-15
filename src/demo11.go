package main

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func max(n1 int, n2 int) int {
	var res int
	if n1 > n2 {
		res = n1
	} else {
		res = n2
	}

	return res

}
func main() {
	v := max(100, 105)
	_, next := swap("a", "b")
	var ptr *string
	// 定义一个空指针
	fmt.Println(ptr)

	var a string = "ptr"
	// ptr指向这个a的地址
	ptr = &a
	fmt.Println(&a)    // a的地址
	fmt.Println(&*ptr) // a的值 获取指针指向的值的地址
	fmt.Println(ptr)   // ptr指向了a的地址
	fmt.Println(&ptr)  // a的地址的地址
	fmt.Println(*ptr)  // a的值
	fmt.Print(v, next)
}
