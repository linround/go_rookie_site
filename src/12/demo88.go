package main

import (
	"../parse"
	"fmt"
)

func main() {
	var example = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"lst class",
		"",
	}
	for _, ex := range example {
		fmt.Printf("解析 %q \n", ex)
		nums, err := parse.Parse(ex)
		fmt.Println("解析出来的值：", nums)
		if err != nil {
			fmt.Println("解析出现错误：", err)
			continue
		}
	}
}
