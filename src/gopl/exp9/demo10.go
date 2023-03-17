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
	response chan<- result // the client wants a single result
}
type entry struct {
	res   result
	ready chan struct{} // res 准备好之后该通道关闭
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (interface{}, error) {
	// 获取缓存的结果
	// 通过直接获取缓存，从而快速的获取已经存在的结果
	memo.mu.Lock()
	e := memo.cache[key]
	// 在新设置之前一直锁住
	// 对于已有的 通过ready通道去传递，从而避免重复的修改，这样也能从通道进行并发的传递
	if e == nil {
		// 对于为缓存
		// 执行函数获取结果，然后放在res下
		e = &entry{ready: make(chan struct{})}
		// res准备好之后，再将res缓存到cache中
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
	} else {
		memo.mu.Lock()
		//
		// 已经有了缓存的话，就需要通过 ready 通道去获取
		// 在没有缓存之前 是直接传输出去的
		<-e.ready // 这个通道 在往外发送。 goroutine必须等待ready之后才能读取条目。即直接从ready channel中读取，而在读取操作关闭之前一直处于阻塞
	}
	// 最终返回缓存结果
	return e.res.value, e.res.err
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
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
