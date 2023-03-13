package main

import "fmt"

func main() {
	months := [...]string{1: "一月", 12: "十二月"}
	//q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(cap(summer), "cap")
	fmt.Println(len(summer), "len")
	// cap是7 不能超过cap
	fmt.Println(summer[:7])
	// summer的长度是3
	endlessSummer := summer[:7]
	fmt.Println(endlessSummer)
	// 修改子串不会影响到底层的
	// 但是修改底层却会影响
	months[12] = "12222"
	// summer0-2对应原来的6-9
	// cap是6-12 所以cap是7
	summer[0] = "88888888888"
	summer[1] = "99999999"
	summer[2] = "7777777"
	fmt.Println(summer[:7])
	fmt.Println(months)
}
