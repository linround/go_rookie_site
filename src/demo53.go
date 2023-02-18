package main

import (
	"bytes"
	"sync"
)

func main() {

}

// RWMutex
// 可以通过 reLock来允许同一时间多个线程对变量进行读操作
// 但是只能一个线程进行写操作
//

type Rw struct {
	rl   sync.RWMutex
	once sync.Once
}

// Once类型 这个方法确保被调用函数只能被调用一次

func tRw(rw *Rw) {
	rw.rl.RLock()
	rw.once.Do(func() {
		//	doSomething
	})
}

// 互斥锁 守护在临界区入口，
// 确保同一时间只能有一个线程进入临界区

type Info struct {
	mu sync.Mutex
}

// 通过mutex来实现一个可以上锁的共享缓冲器件

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// 改变一个变量可以这么做

func Update(info *Info) {
	info.mu.Lock()

	//	doSomething

	info.mu.Unlock()
}
