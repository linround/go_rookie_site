package main

import (
	"fmt"
	"math"
)

// 传统的方式通过函数来计算两个点之间的距离 这个是包级别下的函数

func distance(p, q point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

type point struct {
	x, y float64
}

// p 叫做方法接收器
// 这里定义的是point结构体下声明的方法
func (p point) distance(q point) float64 {

	return math.Hypot(q.x-p.x, q.y-p.y)
}

type path []point

func (pa path) distance() float64 {
	sum := 0.0
	for i := range pa {
		if i > 0 {
			// 计算两个点之间的距离
			sum += pa[i-1].distance(pa[i])
		}
	}
	return sum
}

func main() {
	perim := path{
		{1, 1},
		{1, 2},
		{2, 2},
		{2, 1},
	}
	fmt.Println(perim.distance())
}
