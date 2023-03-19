package main

import (
	"fmt"
	"unsafe"
)

type name1 struct {
	bool
	float64
	int16
}
type name2 struct {
	float64
	int16
	bool
}
type name3 struct {
	bool
	int16
	float64
}

var x struct {
	a bool
	b int16
	c []int
	d name2
}

func main() {
	// Alignof
	// Offsetof
	//Alignof返回对应参数的类型需要对齐的倍数
	fmt.Println("Alignof: x", unsafe.Alignof(x))     // 8
	fmt.Println("Alignof: a", unsafe.Alignof(x.a))   //1
	fmt.Println("Alignof: b", unsafe.Alignof(x.b))   //2
	fmt.Println("Alignof: c", unsafe.Alignof(x.c))   //8
	fmt.Println("Offsetof: a", unsafe.Offsetof(x.a)) //0
	fmt.Println("Offsetof: b", unsafe.Offsetof(x.b)) //2
	fmt.Println("Offsetof: c", unsafe.Offsetof(x.c)) //8 c字段相对于x起始地址的偏移量为8
	fmt.Println("Offsetof: d", unsafe.Offsetof(x.d)) //32 d字段相对于x起始地址的偏移量为32
	fmt.Println("x:",
		unsafe.Sizeof(x))
	fmt.Println("name1", unsafe.Sizeof(name1{false, 2.2, 5}))
	fmt.Println("name2", unsafe.Sizeof(name2{2.2, 5, false}))
	fmt.Println("name3", unsafe.Sizeof(name3{false, 5, 2.2}))

	//            1个字节 = 8 bit
	// 1个机器字 = 8个字节 = 64bit
	fmt.Println(unsafe.Sizeof(float32(0))) // 4个字节
	fmt.Println(unsafe.Sizeof(float64(0))) // 8个字节

	fmt.Println(unsafe.Sizeof(string("")))    // 16个字节
	fmt.Println(unsafe.Sizeof(string("111"))) // 16个字节
	fmt.Println(unsafe.Sizeof(int8(0)))       // 1个字节
	fmt.Println(unsafe.Sizeof(int64(0)))      // 8个字节
	fmt.Println(unsafe.Sizeof(int32(0)))      // 4个字节
	fmt.Println(unsafe.Sizeof(int16(0)))      // 2个字节
	fmt.Println(unsafe.Sizeof(bool(false)))   // 1个字节
}
