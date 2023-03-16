package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//conn, err := net.Dial("tcp", "121.199.1.247:8000") // 远程
	conn, err := net.Dial("tcp", "0.0.0.0:8000") // 本地
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 从连接上读取数据，并写入到控制台 (获取服务器数据)
	go mustCopy(os.Stdout, conn)
	// 从控制台读取数据，并写入连接 (发送给服务器)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// 将写入的流读取到标准输出控制台上
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
