package main

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("出错了")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initSever(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "接收")
		go connectionHandler(conn)
	}

}
func connectionHandler(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	println("连接来自：", connForm)
	syHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		default:
			goto DISCONNECT

		}
	}
DISCONNECT:
	err := conn.Close()
	println("连接关闭")
	checkError(err, "close")

}
func syHello(to net.Conn) {
	obuf := []byte{'l', 'i', 'n', '\'', 'y', 'u', 'a', 'n'}
	wrote, err := to.Write(obuf)
	checkError(err, string(wrote))
}
func initSever(hostAndPort string) net.Listener {
	listener, err := net.Listen("tcp", hostAndPort)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Println("msg", i, msg[i])
		}
		fmt.Printf("接收到的信息：%s", msg)
		print(">")
	}
}
func checkError(err error, info string) {
	if err != nil {
		panic("错误" + info)
	}
}
