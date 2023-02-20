package main

import (
	"bufio"
	"fmt"
	"os"
)

type footer interface {
	foo()
	implementFooter()
}

func main() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.1 / 5212 / go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("请输入姓名")
	// 扫描来自标准输入的文本 将空格分隔的值依次存放到后续的参数内
	fmt.Scanln(&firstName, &lastName)

	fmt.Printf("你好 %s %s", firstName, lastName)
	// 从字符串进行读取
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f, i, s)
	readInput2()
}

func readInput2() {
	// 使用bufio包提供的缓冲读取 来读取数据
	var inputReader *bufio.Reader

	var input string
	var err error
	// 创建一个读取器
	inputReader = bufio.NewReader(os.Stdin)

	fmt.Println("输入一些文字：")

	// ReadString 该方法从输入中读取内容
	// 直到碰到指定字符
	// ReadString 返回读取的字符串，如果碰到错误，则返回
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("你的输入是：", input)

	}
}
