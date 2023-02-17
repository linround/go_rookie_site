package main

import "fmt"

// 类型转换

type INT = int

var z INT = 100

func main() {
	// 具有相同底层类型 int 变量之间的类型转换
	var i int = 100
	c := int(z)
	z = i
	fmt.Println(c)

	var i16 int16 = 10
	var _r int32 = 25
	// 32转换到16
	// 当从一个较大范围转换到较小范围的类型时
	//_r = int32(i16)
	// 16 转为32
	i16 = int16(_r)
	fmt.Println(i16)
}
