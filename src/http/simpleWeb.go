package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = "<h1>title<h1><div>title</div>"

func simpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>HTML模板</h1>")
}
func formServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Println("method:", request.Method)
	switch request.Method {
	case "GET":
		fmt.Println("get 方法")
		io.WriteString(w, form)
	case "POST":
		fmt.Println("post 方法")
		io.WriteString(w, form)
	}
}
func main() {
	http.HandleFunc("/test1", simpleServer)
	http.HandleFunc("/test2", formServer)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("出错了")
	}
}
