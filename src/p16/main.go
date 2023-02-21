package main

import "fmt"

type nextEr interface {
	next() byte
}

func nextFew1(n nextEr, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

// 不要使用一个指针指向接口类型，因为它已经是一个指针
func nextFew2(n *nextEr, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		//b[i] = (n).next // 这里发生编译错误
	}
	return b
}
func main() {
	fmt.Println("开始执行程序")
}
