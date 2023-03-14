package main

import "fmt"

type Buffer struct {
	buf     []byte
	initial [64]byte
}

func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		// 这里发生了内存（cap）分配,
		// 这里初始化了一个长度为0的数组
		b.buf = b.initial[:0]

	}
	//
	if len(b.buf)+n < cap(b.buf) {
		// make
		// 参数1是类型
		// 参数2是长度
		// 参数三是预留内存空间
		buf := make([]byte, len(b.buf)+n, 2*cap(b.buf)+n)
		fmt.Println(len(buf))
		fmt.Println(cap(buf))
		b.buf = buf
	}
	b.buf[5] = 5
	fmt.Println(n)
}

func main() {
	var b Buffer
	b.Grow(20)

}
