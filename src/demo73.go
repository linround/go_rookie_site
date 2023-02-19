package main

import "fmt"

// 不用提前设计出所有的接口
// 可以一个个的接口来进行补充
type shaper interface {
	area() float32
}
type topologicalGenus interface {
	rank() int
}

type square struct {
	side float32
}

func (s *square) area() float32 {
	return s.side * s.side
}
func (s *square) rank() int {
	return 1
}

// 三角形
type rect struct {
	len, wid float32
}

// 实现三角形的rank和area方法
func (r rect) area() float32 {
	return r.wid * r.len
}
func (r rect) rank() int {
	return 2
}

func main() {
	// 本身是值调用
	r := rect{5, 3}
	// 本身是指针调用，
	// 所以对于接口实现的方法
	// 需要保证指针调用
	q := &square{5}
	shapers := []shaper{r, q}
	for n, _ := range shapers {
		fmt.Println(shapers[n])
		fmt.Println(shapers[n].area())
	}

	// 对于新实现的某个接口进行调用
	// 可以看出，接口是可以逐步的去实现的
	// 不必一次性定义好所有的接口
	tgogen := []topologicalGenus{r, q}
	for n, _ := range tgogen {
		fmt.Println(tgogen[n])
		fmt.Println(tgogen[n].rank())
	}
}
