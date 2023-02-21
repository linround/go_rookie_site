package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接错误")
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("你的名字？")
	// 读取控制台的数据 以换行符为结束
	clientName, _ := inputReader.ReadString('\n')

	// 对输入的字符去掉换行和回车
	// 得到客户端输入的名字
	trimmedClient := strings.Trim(clientName, "\r\n")
	//	给服务器发送信息，直到程序退出
	for {
		fmt.Println("发送信息")
		// 再次读取控制台的输入
		input, _ := inputReader.ReadString('\n')
		// 去掉空格和回车
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " 说: " + trimmedInput))
	}
}
