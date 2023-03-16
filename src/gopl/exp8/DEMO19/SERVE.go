package main

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string // 一个即将发送的消息通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // 所有已进入的客户端消息通道
)

func broadcaster() {
	fmt.Println("b1")
	// client 是一个已经接收了消息的通道
	clients := make(map[client]bool)
	fmt.Println("b2")
	for {
		fmt.Println("b3")
		select {
		// 处理messages通道的信息广播
		case msg := <-messages:
			fmt.Println("b4")
			// 将msg发送给所有的客户端
			for cli := range clients {
				fmt.Println("b5")
				cli <- msg
			}
		//	处理进入通道的广播
		// cli时一开始
		case cli := <-entering:
			fmt.Println("b6")
			// 设置已进入即可
			clients[cli] = true
		// 处理离开通道的广播
		case cli := <-leaving:
			fmt.Println("b7")
			// 删除
			delete(clients, cli)
			fmt.Println("b8")
			// 关闭
			close(cli)
			fmt.Println("b9")
		}
	}
}
func handleConn(conn net.Conn) {
	// 发送客户端消息的通道 将军往外发消息的通道
	ch := make(chan string)
	fmt.Println("h1")
	// 所以这里需要在向客户端写入消息前，就创建一个goroutine，进入等待状态
	go clientWriter(conn, ch)
	fmt.Println("h2")
	//
	who := conn.RemoteAddr().String()
	// 将对应数据写入通道
	ch <- "你是 " + who
	fmt.Println("h3")

	// 发送一条消息，并将该数据写入消息通道
	messages <- who + " 已到达"
	fmt.Println("h4")
	// 为进入通道发送一个即将发消息的子通道
	// 除了一开始向客户端写入消息时用到改写入通道
	// 在一开始时，还会将该发送通道传递给enter通道中
	entering <- ch
	// 创建连接的扫描器

	fmt.Println("h5")
	input := bufio.NewScanner(conn)
	// 连接进入时会执行到6
	fmt.Println("h6")
	// 扫描输入
	for input.Scan() {
		fmt.Println("h7")
		// 将对应的信息发送给message
		messages <- who + ": " + input.Text() + "\n"
		fmt.Println("h8")
		// 客户端主动发送消息时执行到71
	}
	fmt.Println("h9")
	// 发送一个即将发消息的通道给leave通道
	leaving <- ch
	fmt.Println("h10")
	// 发送一个已经离开的给消息通道
	messages <- who + " 已离开 \n"
	fmt.Println("h11")
	// 对应连接关闭
	conn.Close()
	fmt.Println("h12")
	//	 一个客户端关闭连接 离开 会执行到11
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// 这一步不太懂
	// 已连接时 这里会向对应的连接客户端发送他的个人连接信息
	fmt.Println("c1")
	for msg := range ch {
		fmt.Println("c2")
		fmt.Fprintf(conn, msg)
		fmt.Println("c3")
	}
}

func main() {
	fmt.Println("m1")
	// 开启一个监听服务 在8000端口
	listener, _ := net.Listen("tcp", "0.0.0.0:8000")
	fmt.Println("m2")
	// 这里需要先开启一个广播 goroutine
	// 监听消息广播，进入广播 ，离开广播
	go broadcaster()
	fmt.Println("m3")
	for {
		fmt.Println("m4")
		// 只执行serve时会执行到m4
		conn, _ := listener.Accept()
		// (当有一个连接进入时会执行到m5,m6,1,2,3,4,5,6,m4)再次进入到m4的阻塞等待
		fmt.Println("m5")
		// 这里开启了一个连接的goroutine
		// 当没有连接是这个for循环处于阻塞
		// 当有一个链接时会开启
		go handleConn(conn)
		fmt.Println("m6")
	}
}

// 启动服务的过程
// m1 m2 m3 m4 b1 b2 b3

// 连接接入
// m5 m6 m4 h1 h2 c1 h3 h4 c2 b4 b3 b6 b3 h5 h6 c3

// 发送第一条消息
// h7 h8 b4 b5 b3 c2 c3

// 退出
// h9 h10 b7 b8 b9 b3 b4 b3 h11 h12
