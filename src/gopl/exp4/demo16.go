package main

import "fmt"

type point struct {
	X, Y int
}

type circle struct {
	//嵌入式结构体的命名
	point
	//重复属性的访问
	X      int
	Radius int
}

type wheel struct {
	circle
	Spokes int
}

// 嵌入式结构体的命名、属性访问、重复属性的访问
func main() {
	var w wheel
	var t wheel = wheel{
		Spokes: 100,
		circle: circle{Radius: 122, X: 78, point: point{X: 52, Y: 78}},
	}
	fmt.Printf("%v", t)
	fmt.Printf("%v", t.X)
	fmt.Printf("%v", t.circle.point.X)
	w.circle.point.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

}
