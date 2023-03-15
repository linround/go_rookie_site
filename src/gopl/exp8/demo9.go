package main

import "fmt"

// 会与goroutine的应用

// 第一个goroutine生成数字
// 第二个goroutine计算数字平方
// 第三个goroutine 打印计算结果

// 所以需要两个channel

// 以下演示了逐个关闭channel
// 从而实现打印指定的数字

func main() {
	nat := make(chan int)
	squ := make(chan int)

	//	输出数
	go func() {
		for x := 0; x <= 10; x++ {
			// 通过通道nat传递数
			nat <- x
		}
		close(nat)
	}()
	//	计算平方
	go func() {
		for {
			// 接收数字
			x, ok := <-nat
			if !ok {
				break
			}
			// 传递平方（如果上级被关闭 那么就会一直穿0）
			squ <- x * x
		}
		close(squ)
	}()
	// 无限等待
	for {
		v, ok := <-squ
		if !ok {
			break
		}
		fmt.Println(v)
	}

	//以下结合time.Sleep进行程序等待
	//go func() {
	//	// 等待接收平方值
	//	for {
	//		fmt.Println(<-squ)
	//	}
	//}()
	//time.Sleep(5 * time.Second)

}
