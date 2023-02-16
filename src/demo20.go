package main

import "fmt"

// 定义一个接口

type Phone interface {
	// 定义一个方法及其返回类型

	test()
	call()
}

//定义一个结构体

type NokiaPhone struct {
}

// 实现结构体中的方法

func (t NokiaPhone) test() {
	fmt.Println("结构体nokiaPhone test方法")
}
func (t NokiaPhone) call() {
	fmt.Println("结构体NokiaPhone call方法")
}

// 定义一个结构体

type IPhone struct {
}

// 实现这个结构体中的方法
func (i IPhone) test() {
	fmt.Println("结构体 IPhone  test方法")
}
func (i IPhone) call() {
	fmt.Println("结构体 IPhone call方法")
}

func main() {
	// 定义一个phone 变量  该变量类型是Phone 接口
	//
	// 定义一个接口 这个接口 有需要实现的方法  这个接口可以当作类型
	// 使用接口指定变量的类型

	// 定义一个结构体积  该结构提可以通过 new 来构造对象
	// 实现对应结构体中的方法

	//最终有一个 Phone 接口类型的变量
	// 该变量被赋值一个对象
	// 该对象是通过new一个结构体出来的
	// 该结构体正好实现了该接口中的方法

	var phone Phone
	phone = new(NokiaPhone)
	phone.test()
	phone.call()

	phone = new(IPhone)
	phone.test()
	phone.call()
}
