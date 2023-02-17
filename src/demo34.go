package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "l english li"
	index := strings.Index(s, "n")
	println(index)
	sr := strings.Replace(s, "l", "oo", -1)
	println(sr)
	tri := strings.Trim(s, "l")
	println(tri)
	trs := strings.TrimSpace(tri)
	println(trs)
	sli := strings.Fields(trs)
	println(sli)
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	println(sl)

	for _, v := range sl {
		fmt.Printf("%s -", v)
	}
	println()
	str2 := "l|d|p|o"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("%v \n", sl2)

}
