package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
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

func echo(c net.Conn, shout string, delay time.Duration) {
	// 为这个连接写入 对应的大写字符（发送大写数据）
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	// 为这个连接写入 对应的字符（发送正常数据）
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	// 为这个连接写入 小写的字符（发送小写数据）
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		// 为了防止客户端的第三次shout在前一个shout处理完成之前一直没有被处理
		// 使用go关键字开启协程
		//
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
