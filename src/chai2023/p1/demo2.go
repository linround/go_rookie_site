package main

import "fmt"

func main() {
	f1()

}
func f1() {
	x := []int{2, 3, 5, 7, 11}
	y := x[1:3]
	fmt.Println("x：", len(x), cap(x))
	fmt.Println("y：", len(y), cap(y))

	var (
		a []int // nil 切片, 和 nil 相等, 一般用来表示一个不存在的切片

		b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有 3 个元素的切片, len 和 cap 都为 3
		d = c[:0]             // 有 2 个元素的切片, len 为 2, cap 为 3
		e = c[0:2:cap(c)]     // 有 2 个元素的切片, len 为 2, cap 为 3
		f = c[:0]             // 有 0 个元素的切片, len 为 0, cap 为 3
		g = make([]int, 3)    // 有 3 个元素的切片, len 和 cap 都为 3
		h = make([]int, 2, 3) // 有 2 个元素的切片, len 为 2, cap 为 3
		i = make([]int, 0, 3) // 有 0 个元素的切片, len 为 0, cap 为 3
	)
	g = append(g, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	g = append(g[:1], append([]int{2}, g[1:]...)...)
	fmt.Println(g)
	fmt.Println("a：", len(a), cap(a))
	fmt.Println("b：", len(b), cap(b))
	fmt.Println("c：", len(c), cap(c))
	fmt.Println("d：", len(d), cap(d))
	fmt.Println("e：", len(e), cap(e))
	fmt.Println("f：", len(f), cap(f))
	fmt.Println("g：", len(g), cap(g))
	fmt.Println("h：", len(h), cap(h))
	fmt.Println("i：", len(i), cap(i))

}
