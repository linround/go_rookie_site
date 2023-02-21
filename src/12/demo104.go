package main

import (
	"fmt"
)

//futures模式 指的是你在使用某一个值之前需要先对其进行计算
// 这种情况下，就可以在另一个处理器上进行该值的计算，在使用时，
// 该值就已经计算完毕了

// 假设我们有一个矩阵类型，我们需要计算两个矩阵A和B乘积的逆，
// 首先我们通过函数 Inverse(M) 分别对其进行求逆运算，在将结果相乘。

// 在这个例子中，a和b的求逆矩阵需要先被计算。那么为什么在计算b的逆矩阵时，
// 需要等待a的逆计算完成呢？显然不必要，这两个求逆运算其实可以并行执行的。

// 求两个矩阵乘积
func inverseProduct(a int, b int) int {
	aFuture := inverseFuture(a)
	bFuture := inverseFuture(b)
	aInv := <-aFuture
	bInv := <-bFuture
	return pr(aInv, bInv)

}
func pr(a, b int) int {
	return a * b
}

// 求一个矩阵的乘积
func inverseFuture(a int) chan int {
	future := make(chan int)
	go func() {
		future <- 5 * a
	}()
	return future
}

func main() {
	res := inverseProduct(4, 5)
	fmt.Println("res", res)
}
