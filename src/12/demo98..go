package main

import "fmt"

// 发送数据
func generate1() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 50; i++ {
			ch <- i
		}
	}()
	return ch
}

// 通过prime来过滤输入值，然后发送给返回通道
func filter1(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate1()
		for {
			prime := <-ch
			fmt.Println("prime", prime)
			ch = filter1(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		// 进入无限等待
		// 等待协程有输入后即可直接获取
		<-primes
	}
}
