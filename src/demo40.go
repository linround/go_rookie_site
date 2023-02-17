package main

import "fmt"

func main() {
	res := sw(2)
	println(res)
	println("fo开始")
	fo()
}

func fo() {
	// 中文编码乱码
	str := "日本語"
	println(str)
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("%c", str[ix])
	}
	str2 := "中文"
	println(str2)
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("%c \n", str2[ix])
	}
	str3 := "english"
	println(str3)
	for ix := 0; ix < len(str3); ix++ {
		fmt.Printf("%c \n", str3[ix])
	}
}

func sw(i int) int {
	switch i {
	case 0:
		println(0)
		return 0
	case 1:
		println(1)
		return 1
	case 2:
		println(2)
		// 此处return 之后，不会因为 fallthrough的存在，继续执行下一步
		// 可使用return语句来提前结束分支
		//return 2
		fallthrough
	default:
		println("end")
		return 100
	}
}
