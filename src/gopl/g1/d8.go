package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // 这是一个互斥锁
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	// 然而在并发情况下，假如真的有两个请求同一时刻去更新count，
	//那么这个值可能并不会被正确地增加；这个程序可能会引发一个严重的bug：竞态条件。
	//为了避免这个问题，我们必须保证每次修改变量的最多只能有一个goroutine，
	//这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "path = %q\n", r.URL.Path)
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d \n", count)
	mu.Unlock()
}
