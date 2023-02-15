package main

import (
	"fmt"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	// Sprintf 生成格式化字符串 并返回该字符串
	// Printf 生成格式化字符串并写入标准输出
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)
}
