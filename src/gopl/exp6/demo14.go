package main

import (
	"fmt"
	"io"
	"os"
)

type a interface {
}

func main() {
	var a1 = a("success1")
	// a 是一个接口类型
	// a1.(a) 用于判断 a1 是否是a这个接口类型
	// 最终输出 值和判断结果
	f, ok := a1.(a)
	fmt.Println("a1: ", f, ok) // success1 true
	m1()

}

func m1() {
	var w io.Writer
	fmt.Println(w == nil) // true
	// 断言操作的对象是一个nil接口值，那么不论被断言的类型是什么这个类型断言都会失败
	w, ok0 := w.(io.ReadWriter) // false
	fmt.Println("f0:", w, ok0)
	var r io.Reader
	w = os.Stdout
	r = os.Stdout
	// 以下可以理解为：
	// Writer 是一个单独的接口类型
	// ReadWriter 相当于继承有Reader、Writer两个接口类型
	// 所以 ReadWriter 也相当于是一个Writer类型
	w, ok := w.(io.ReadWriter) // success: f == os.Stdout
	fmt.Println("f:", w, ok)
	f2, ok2 := r.(io.ReadWriter) // success: f == os.Stdout
	fmt.Println("f2:", f2, ok2)
	// w和rw都持有os.Stdout，因此它们都有一个动态类型*os.File
	f3, ok3 := r.(*os.File)
	fmt.Println("f3:", f3, ok3) // success: f == os.Stdout

}
