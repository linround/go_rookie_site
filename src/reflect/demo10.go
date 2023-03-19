package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
	d int64
}

func main() {

	// 和 pb := &x.b 等价
	// x的起始地址 + b相对于x的偏移地址  得到b的地址
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	fmt.Println(&(x).b) // 地址 0x9bff42
	fmt.Println(pb)     // 地址 0x9bff42
	fmt.Println(*pb)
	*pb = 42
	fmt.Println(x.b) // "42"

	// 当一个变量被移动时
	// 所有保存改变量旧地址的指针必须 同时被更新为变量移动后的新地址
	// 一个unsafe.Pointer是一个指向变量的指针，因此当变量被移动时对应的指针也必须被更新
	// 但是uintptr类型的临时变量只是一个普通的数字，所以其值不应该被改变
	//  面错误的代码因为引入一个非指针的临时变量tmp，导致垃圾收集器无法正确识别这个是一个指向变量x的指针
	// 当第二个语句执行时，变量x可能已经被转移，这时候临时变量tmp也就不再是现在的&x.b地址。第三个向之前无效地址空间的赋值语句将彻底摧毁整个程序

	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b) // 尽量少用中间值来取地址
	fmt.Println("temp:", tmp)
	pb = (*int16)(unsafe.Pointer(tmp))
	*pb = 46
	fmt.Println(x.b)

	// 这里并没有指针引用new新创建的变量，因此该语句执行完成之后，垃圾收集器有权马上回收其内存空间，所以返回的pT将是无效的地址

	pT := uintptr(unsafe.Pointer(new(T))) // 提示: 错误!

}
