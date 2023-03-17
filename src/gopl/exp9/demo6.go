package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex // guards balance

type test struct {
	name string
}

func (t *test) getName(s int) string {
	fmt.Println(" -----开始加锁", s)
	mu.RLock()
	fmt.Println(" 已经加锁", s)
	t.name = ";;;;;;;;;"
	name := t.name
	fmt.Println(" 已经赋值", s)
	//defer mu.RUnlock()
	return name
}

// 创建一个加法函数试试
// create
var vt test = test{"linyuan"}

func main() {
	go func() {
		fmt.Println("打印值111:", vt.getName(111))
	}()
	go func() {
		fmt.Println("打印值222:", vt.getName(222))
	}()
	go func() {
		fmt.Println("打印值333:", vt.getName(333))
	}()
	go func() {
		fmt.Println("打印值444:", vt.getName(444))
	}()
	go func() {
		fmt.Println("打印值555:", vt.getName(555))
	}()
	time.Sleep(3 * time.Second)
}
