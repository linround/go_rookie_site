package main

// 取地址符 &号放到变量前就会返回相应变量的地址
// *号用于指定变量作为一个指针
func main() {
	// 定义一个整型指针
	var ptr *int
	println(ptr)
	var a int = 1000
	// 将这个整型指针指向变量a的地址
	println(&a)
	ptr = &a
	// 打印这个地址
	println(ptr)
	println()
	// 打印这个指针指向的值
	println(*ptr)

}
