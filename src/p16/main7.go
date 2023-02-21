package main

import "log"

// 关于使用recover 终止panic过程

func protect(g func()) {
	defer func() {
		log.Println("完成")
		if x := recover(); x != nil {
			log.Println("出现panic", x)
		}
	}()
	log.Println("开始")
	g()
}
func pa() {
	println("panic之前")
	panic("发送panic")
	// 这里没有被保护所以以下程序不执行
	println("panic之后继续执行")
}
func main() {
	// 此处使用protect来模仿catch
	protect(pa)
	// 即便上述程序出现了panic
	// 但是已经被捕获，所以还是会继续执行
	println("main")
}
