package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

func (memo *Memo) Get(key string) (interface{}, error) {
	// 获取缓存的结果
	// 通过直接获取缓存，从而快速的获取已经存在的结果
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	// 这里存在的问题
	// 有一些Url被获取两次
	// 多个goroutine一起查询cache，发现没有值
	// 然后一起调用f这个慢不拉叽的函数
	// 在得到结果后，也都会去更新map。其中一个获得的结果会覆盖掉另一个的结果

	// 理想情况应该是避免重复修改（重复抑制/避免）
	if !ok {
		// 对于为缓存
		// 执行函数获取结果，然后放在res下
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		// res准备好之后，再将res缓存到cache中
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	// 最终返回缓存结果
	return res.value, res.err
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func main() {
	m1()
}

var n sync.WaitGroup

func m1() {
	m := New(httpGetBody)
	var urls = []string{
		"https://pkg.go.dev/",
		"http://127.0.0.1:9527",
		"https://pkg.go.dev/",
		"https://pkg.go.dev/",
		"https://pkg.go.dev/",
		"https://pkg.go.dev/"}
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				fmt.Println(url, ":", err)
			}
			fmt.Println(url, "=", time.Since(start), "字节：", len(value.([]byte)))
			// n.Done只负责减一，
			// 不会阻塞下一个
			n.Done()
		}(url)
	}
	// n 在0的时候才会继续执行n.wait的位置
	n.Wait()
	fmt.Println("完成")

}
