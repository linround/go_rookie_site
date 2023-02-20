package main

import (
	"fmt"
	"os"
	"strings"
)

//go run demo79.go param
//demo79
//你好 linyuanparam

// 在运行脚本时可以添加相关的一些参数
func main() {
	fmt.Println("demo79")
	who := "linyuan"
	if len(os.Args) > 0 {
		who += strings.Join(os.Args[1:], "")
	}
	fmt.Println("你好", who)

}
