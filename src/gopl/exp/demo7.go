package main

import "fmt"

func main() {
	fmt.Println(5.0 / 4.0)
	fmt.Println(5 / 4)
	fmt.Println(5.0 / 4)
	fmt.Println(5 / 4.0)
	fmt.Printf("%b", -5)
	fmt.Println()
	fmt.Println(-2 << 1)
	var p = new(int)
	fmt.Println(*p)
	f := 257.141
	i := uint8(f)
	fmt.Println(i)
	//	字符使用c% 参数打印；使用%q参数打印带单引号的字符
	single := '4'
	fmt.Printf("%c \n", single)
	fmt.Printf("%q \n", single)
}
