package main

import "fmt"

func main() {
	fun1()
	fun3()
	fun4()
	v := new(int8)
	println()
	fmt.Printf("vvvvvvvvvv %d", v)
	println()
}

// defer 允许我们推迟到函数返回之前一刻才执行某个语句或函数
// 一般用于释放已分配的资源
func fun1() {
	defer fun2()
	fmt.Print("f1 top \n")
	fmt.Print("f1 bottom \n")
	return

}
func fun2() {
	fmt.Print("f2 被执行 \n")
}

func fun3() {
	defer P(82)
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func P(i int) {
	fmt.Printf("f3 被执行 %d \n", i)
}

func fun4() {
	println("f4执行：")
	defer fmt.Print(522222)
	i := 0
	defer fmt.Println(i)
	i++
	return
}
