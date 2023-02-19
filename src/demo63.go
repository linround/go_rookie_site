package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

// 定义了一个接收者类型是Square的area方法

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	// 创建了一个Square实例
	sq1 := new(Square)
	sq1.side = 5
	// 所以可以将一个Square 类型的变量赋值给
	// 一个接口类型的变量
	var areaInt Shaper
	// 现在接口变量包含一个指向Square变量的引用，
	// 从而可以调用Square上的方法Area
	// 在接口实例上调用，可以使得此方法更具有一般性
	areaInt = sq1
	fmt.Println(areaInt.Area())
}
