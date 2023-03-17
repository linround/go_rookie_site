package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type request struct {
	key      string
	response chan<- result // response是一个接收通道，接收的类型为result
}
type entry struct {
	res   result
	ready chan struct{} // res 准备好之后该通道关闭
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}
func (e *entry) deliver(response chan<- result) {
	<-e.ready
	// 接收通道 接收res
	response <- e.res
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	requests chan request
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response // response 将接收到的数据发送出来（response的功能是从别处接收，然后再发给别人，其是接收通道）
	return res.value, res.err
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
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
	// 对于同一个url 多个goroutine 等待的时间是一样的，只需要从同一个通道中接收该信息即可
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
