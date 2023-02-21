package main

import (
	"fmt"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockCode = 123
	var endDate = "2020-p12-31"
	var url = "Code=%d&endDate=%s"
	fmt.Printf(url, stockCode, endDate)
}
