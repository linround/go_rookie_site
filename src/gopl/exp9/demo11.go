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
	// 收到了数据，这个通道可以关闭了，从而防止继续阻塞
	fmt.Println("call1")
	e.res.value, e.res.err = f(key)
	// 如不关闭ready通道,则会一直阻塞
	fmt.Println("call2")
	close(e.ready)
	fmt.Println("call3")
}
func (e *entry) deliver(response chan<- result) {
	fmt.Println(<-e.ready, "deliver1")
	// 接收通道 接收res
	response <- e.res
	fmt.Println(<-e.ready, "deliver2")
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

	fmt.Println("s1")
	cache := make(map[string]*entry)
	fmt.Println("s2")
	for req := range memo.requests {
		fmt.Println("s3")
		e := cache[req.key]
		fmt.Println("s4")
		if e == nil {
			fmt.Println("s5")
			e = &entry{ready: make(chan struct{})}
			fmt.Println("s6")
			cache[req.key] = e
			fmt.Println("s7")
			// 在这里会先进入一个阻塞
			// 在这个goroutine中，会等待对e的赋值操作
			go e.call(f, req.key)
			fmt.Println("s8")
		}
		fmt.Println("s9")
		// 对于已经赋值的操作 构建一个并发
		// 传入这个req的response通达，将值通过response通道传递出去
		go e.deliver(req.response)
		fmt.Println("s10")
	}
}

func New(f Func) *Memo {
	fmt.Println("n1")
	memo := &Memo{requests: make(chan request)}
	fmt.Println("n2")
	go memo.server(f)
	fmt.Println("n3")
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

func (memo *Memo) Get(key string) (interface{}, error) {
	fmt.Println("g1", key)
	response := make(chan result)
	fmt.Println("g2", key)
	// requests 是一个传递request类型的通道
	// 在这里将对应的url和其响应接收器 传入requests通道
	memo.requests <- request{key, response}
	fmt.Println("g3", key)
	// 这个响应通道还未接收到信息之前是阻塞的
	// 接收到信息后立即将信息传递到res
	res := <-response // response 将接收到的数据发送出来（response的功能是从别处接收，然后再发给别人，其是接收通道）
	fmt.Println("g4", key)
	return res.value, res.err
}
func m1() {
	m := New(httpGetBody)
	// 对于同一个url 多个goroutine 等待的时间是一样的，只需要从同一个通道中接收该信息即可
	var urls = []string{
		//"https://pkg.go.dev/",
		//"http://127.0.0.1:9527",
		//"https://pkg.go.dev/",
		"https://pkg.go.dev/",
		//"https://pkg.go.dev/",
		"https://pkg.go.dev/"}
	for _, url := range urls {
		fmt.Println("mFor", url)
		// 创建多个goroutine,然后执行到wait状态进入等待
		// 这里创建的goroutine 进入并发状态，无法预知顺序
		go func(url string) {
			fmt.Println("m1", url)
			start := time.Now()
			fmt.Println("m2", url)
			value, err := m.Get(url)
			fmt.Println("m3", url)
			if err != nil {
				fmt.Println(url, ":", err)
			}
			fmt.Println(url, "=", time.Since(start), "字节：", len(value.([]byte)))
			// n.Done只负责减一，
			// 不会阻塞下一个
			fmt.Println("m4", url)
			fmt.Println("m5", url)
		}(url)
	}
	fmt.Println("mWait")
	// n 在0的时候才会继续执行n.wait的位置
	time.Sleep(5 * time.Second)
	fmt.Println("完成")

}
