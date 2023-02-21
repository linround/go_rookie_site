package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.baidu.com"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	//	 创建socket连接
	con, err := net.Dial("tcp", remote)
	// 发送一个消息
	io.WriteString(con, msg)
	fmt.Println("读取响应")
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Println("read", read)
		fmt.Println(string(data[0:count]))
	}
	con.Close()
}
