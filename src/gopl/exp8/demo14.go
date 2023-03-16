package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

func main() {
	worklist := make(chan []string)
	//myUrl := []string{"https://pkg.go.dev/html/template"}
	urls := []string{"http://121.199.1.247/", "https://golang.org/help/", "https://golang.org/doc/", "https://golang.org/blog/"}
	// Start with the command-line arguments.
	go func() { worklist <- urls }()
	var level = 3
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	// 这里如果 持续的传入，会导致无限的递归并发
	for list := range worklist {
		fmt.Println(level)
		if level == 0 {
			close(worklist)
			break
		}
		level--
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					// 广度优先的方式进行继续抓取
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// 并发的web爬虫
func crawl(url string) []string {
	// 抓取网站中的url链接
	list, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list
}
