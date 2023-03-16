package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

var tokens = make(chan struct{}, 20)

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	var urls = []string{"https://github.com/linround"}
	go func() { worklist <- urls }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		// list从worklist channel接收后
		list := <-worklist
		// 继续遍历对应的链接
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				//使用 类似信号量的n值 也会导致相关的无限并发
				n++
				// 这里会导致 根据links的数目去不断开启协程
				// 可能会造成开始上万个协程的性能问题
				go func(link string) {
					// 抓取到的links通过worklist channel传递到list,
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
