package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("服务端处理····")

	// 在response中写入要传输的数据
	// 这里时写入了一个字符串
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", "title", "body")

}
func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		log.Fatal("创建服务失败····")
	}
}
