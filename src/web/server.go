package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务启动")
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("启动程序错误")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接错误")
			return
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		// 使用一个512字节的缓冲
		buf := make([]byte, 1024*500)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取粗偶")
			return
		}
		fmt.Printf("收到%d数据 %v \n", len, string(buf[:len+50]))
	}
}
