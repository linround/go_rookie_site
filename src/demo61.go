package main

import "fmt"

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return p.x * p.y
}

type NamedPoint struct {
	Point
	name string
}

// 子的Abs只作用在子上
// 通过重写Abs 来改变

func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

func main() {
	n := new(NamedPoint)
	n.x = 3
	n.y = 4
	n.name = "PY"
	fmt.Println(n.Abs())
}
