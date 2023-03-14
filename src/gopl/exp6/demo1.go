package main

import (
	"fmt"
	"image/color"
	"math"
)

type point struct {
	x, y float64
}

func (p *point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p *point) distance(q point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

type colorPoint struct {
	point
	color color.RGBA
}

func main() {
	var cp colorPoint
	cp.x = 1
	fmt.Println(cp.point.x)
	cp.point.y = 2
	fmt.Println(cp.y)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = colorPoint{point{1, 1}, red}
	var q = colorPoint{point{5, 4}, blue}
	fmt.Println(p.distance(q.point))
	p.scaleBy(2)
	q.scaleBy(2)

	fmt.Println(p.distance(q.point))
}
