package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	f1()
	f2()
	f3()
	f5()

}
func f2() {
	for i, c := range []byte("世界") {
		fmt.Println("字节码", i, c)
	}
}
func f1() {
	s := "a世界"
	for i, c := range s {
		fmt.Println("原始", i, c, s[i])
	}
}
func f3() {
	var v = []byte{228, 184, 150}
	fmt.Println("真实编码", string(v))
}
func f5() {
	fmt.Printf("run: %#v\n", string([]rune{19990, 30028}))
}

func f4() {

	c1 := make(chan int)
	go func() {
		fmt.Println("c1")
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- 100 * 85
		}()
	}()
	fmt.Println(unsafe.Sizeof("你好"))
	// 通道等待
	fmt.Println(<-c1)
	fmt.Printf("%#v\n", []byte("世界"))
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8cabc") // �界abc

}
