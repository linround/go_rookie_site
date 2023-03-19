package main

import "fmt"

func main() {
	f1()
}
func f1() {
	x := 1
	y := 2
	var a = []*int{&x, &y}
	fmt.Println(a[len(a)-1])
	a = a[:len(a)-1] // 被删除的最后一个元素依然被引用, 可能导致 GC 操作被阻碍
	fmt.Println(a[len(a)-1])
}
