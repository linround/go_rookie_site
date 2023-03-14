package main

import "fmt"

// 推理顺序
// m1 t1 t2 t m2 m

func test() {
	defer func() {
		fmt.Println("执行完毕test")
	}()
	fmt.Println("test1")
	fmt.Println("test2")
}
func main() {
	defer func() {
		fmt.Println("执行完毕main")
	}()
	fmt.Println("main1")
	test()
	//panic("SSS")
	fmt.Println("main2")
}
