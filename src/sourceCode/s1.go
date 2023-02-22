package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//duration()
	str()
}

func str() {
	var s1 string = "   "
	s2 := "  你好"
	fmt.Println(strings.TrimSpace(s1), "=s1")
	fmt.Println(strings.TrimSpace(s2), "=s2")
	s3 := 7
	td := time.Duration(s3)
	fmt.Println(td)
}
func duration() {
	hour, _ := time.ParseDuration("10h")
	fmt.Println(hour)
	day, _ := time.ParseDuration("7d")
	fmt.Println(day)
}
