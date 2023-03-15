package main

import (
	"io"
	"net"
	"time"
)

// 每隔一秒 将当前时间写到客户端

func main() {
	// 创建一个lis对象
	// 这个对象会监听一个网络端口上的到来的链接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		return
	}
	for true {

		// listener.Accept对象会直接阻塞，直到一个新的连接被创建

		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 处理一个完整的客户端连接
		// 增加go关键字 每一次连接都会进入一个独立的goroutine
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// 获取到当前时刻，然后写入到客户端
		_, err := io.WriteString(c, time.Now().Format("15:04:05:00  linyuan \n"))
		if err != nil {
			return
		}
		// 每隔疫苗及就执行写入
		time.Sleep(1 * time.Second)
	}
}
