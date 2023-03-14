package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
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
	*point
	color color.RGBA
}

func main() {
	//var cp colorPoint
	//cp.point.x = 1
	//fmt.Println(cp.point.x)
	//cp.point.y = 2
	//fmt.Println(cp.y)
	var cache = struct {
		sync.Mutex
		mapping map[string]string
	}{
		mapping: map[string]string{"name": "testName"},
	}
	cache.Lock()
	fmt.Println(cache.mapping["name"])
	cache.Unlock()
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = colorPoint{&point{1, 1}, red}
	var q = colorPoint{&point{5, 4}, blue}
	fmt.Println(p.distance(*q.point))
	p.scaleBy(2)
	q.scaleBy(2)

	fmt.Println(p.distance(*q.point))
}
